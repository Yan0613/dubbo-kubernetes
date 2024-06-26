/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mux

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net/url"
	"os"
	"time"
)

import (
	"github.com/go-logr/logr"

	"github.com/pkg/errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

import (
	mesh_proto "github.com/apache/dubbo-kubernetes/api/mesh/v1alpha1"
	"github.com/apache/dubbo-kubernetes/pkg/config/multizone"
	"github.com/apache/dubbo-kubernetes/pkg/core"
	"github.com/apache/dubbo-kubernetes/pkg/core/runtime/component"
	"github.com/apache/dubbo-kubernetes/pkg/dds"
	"github.com/apache/dubbo-kubernetes/pkg/version"
)

const (
	DDSVersionHeaderKey = "dds-version"
	DDSVersionV3        = "v3"
)

var muxClientLog = core.Log.WithName("dds-mux-client")

type client struct {
	globalToZoneCb OnGlobalToZoneSyncStartedFunc
	zoneToGlobalCb OnZoneToGlobalSyncStartedFunc
	globalURL      string
	clientID       string
	config         multizone.DdsClientConfig
	ctx            context.Context
}

func NewClient(
	ctx context.Context,
	globalURL string,
	clientID string,
	globalToZoneCb OnGlobalToZoneSyncStartedFunc,
	zoneToGlobalCb OnZoneToGlobalSyncStartedFunc,
	config multizone.DdsClientConfig,
) component.Component {
	return &client{
		ctx:            ctx,
		globalToZoneCb: globalToZoneCb,
		zoneToGlobalCb: zoneToGlobalCb,
		globalURL:      globalURL,
		clientID:       clientID,
		config:         config,
	}
}

func (c *client) Start(stop <-chan struct{}) (errs error) {
	u, err := url.Parse(c.globalURL)
	if errs != nil {
		return err
	}
	dialOpts := []grpc.DialOption{}
	dialOpts = append(dialOpts, grpc.WithUserAgent(version.Build.UserAgent("dds")), grpc.WithDefaultCallOptions(
		grpc.MaxCallSendMsgSize(int(c.config.MaxMsgSize)),
		grpc.MaxCallRecvMsgSize(int(c.config.MaxMsgSize))),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                grpcKeepAliveTime,
			Timeout:             grpcKeepAliveTime,
			PermitWithoutStream: true,
		}),
	)
	switch u.Scheme {
	case "grpc":
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	case "grpcs":
		tlsConfig, err := tlsConfig(c.config.RootCAFile, c.config.TlsSkipVerify)
		if err != nil {
			return errors.Wrap(err, "could not ")
		}
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	default:
		return errors.Errorf("unsupported scheme %q. Use one of %s", u.Scheme, []string{"grpc", "grpcs"})
	}
	conn, err := grpc.Dial(u.Host, dialOpts...)
	if err != nil {
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			errs = errors.Wrapf(err, "failed to close a connection")
		}
	}()
	withDDSCtx, cancel := context.WithCancel(metadata.AppendToOutgoingContext(c.ctx,
		"client-id", c.clientID,
		DDSVersionHeaderKey, DDSVersionV3,
		dds.FeaturesMetadataKey, dds.FeatureZonePingHealth,
		dds.FeaturesMetadataKey, dds.FeatureHashSuffix,
	))
	defer cancel()

	log := muxClientLog.WithValues("client-id", c.clientID)
	errorCh := make(chan error)

	go c.startHealthCheck(withDDSCtx, log, conn, errorCh)

	go c.startGlobalToZoneSync(withDDSCtx, log, conn, errorCh)
	go c.startZoneToGlobalSync(withDDSCtx, log, conn, errorCh)

	select {
	case <-stop:
		cancel()
		return errs
	case err = <-errorCh:
		cancel()
		return err
	}
}

func (c *client) startGlobalToZoneSync(ctx context.Context, log logr.Logger, conn *grpc.ClientConn, errorCh chan error) {
	ddsClient := mesh_proto.NewDDSSyncServiceClient(conn)
	log = log.WithValues("rpc", "global-to-zone")
	log.Info("initializing Dubbo Discovery Service (DDS) stream for global to zone sync of resources with delta xDS")
	stream, err := ddsClient.GlobalToZoneSync(ctx)
	if err != nil {
		errorCh <- err
		return
	}
	processingErrorsCh := make(chan error)
	c.globalToZoneCb.OnGlobalToZoneSyncStarted(stream, processingErrorsCh)
	c.handleProcessingErrors(stream, log, processingErrorsCh, errorCh)
}

func (c *client) startZoneToGlobalSync(ctx context.Context, log logr.Logger, conn *grpc.ClientConn, errorCh chan error) {
	ddsClient := mesh_proto.NewDDSSyncServiceClient(conn)
	log = log.WithValues("rpc", "zone-to-global")
	log.Info("initializing Dubbo Discovery Service (DDS) stream for zone to global sync of resources with delta xDS")
	stream, err := ddsClient.ZoneToGlobalSync(ctx)
	if err != nil {
		errorCh <- err
		return
	}
	processingErrorsCh := make(chan error)
	c.zoneToGlobalCb.OnZoneToGlobalSyncStarted(stream, processingErrorsCh)
	c.handleProcessingErrors(stream, log, processingErrorsCh, errorCh)
}

func (c *client) startHealthCheck(
	ctx context.Context,
	log logr.Logger,
	conn *grpc.ClientConn,
	errorCh chan error,
) {
	client := mesh_proto.NewGlobalDDSServiceClient(conn)
	log = log.WithValues("rpc", "healthcheck")
	log.Info("starting")

	prevInterval := 5 * time.Minute
	ticker := time.NewTicker(prevInterval)
	defer ticker.Stop()
	for {
		log.Info("sending health check")
		resp, err := client.HealthCheck(ctx, &mesh_proto.ZoneHealthCheckRequest{})
		if err != nil && !errors.Is(err, context.Canceled) {
			if status.Code(err) == codes.Unimplemented {
				log.Info("health check unimplemented in server, stopping")
				return
			}
			log.Error(err, "health check failed")
			errorCh <- errors.Wrap(err, "zone health check request failed")
		} else if interval := resp.Interval.AsDuration(); interval > 0 {
			if prevInterval != interval {
				prevInterval = interval
				log.Info("Global CP requested new healthcheck interval", "interval", interval)
			}
			ticker.Reset(interval)
		}

		select {
		case <-ticker.C:
			continue
		case <-ctx.Done():
			log.Info("stopping")
			return
		}
	}
}

func (c *client) handleProcessingErrors(
	stream grpc.ClientStream,
	log logr.Logger,
	processingErrorsCh chan error,
	errorCh chan error,
) {
	err := <-processingErrorsCh
	if status.Code(err) == codes.Unimplemented {
		log.Error(err, "rpc stream failed, because global CP does not implement this rpc. Upgrade remote CP.")
		// backwards compatibility. Do not rethrow error, so DDS multiplex can still operate.
		return
	}
	if errors.Is(err, context.Canceled) {
		log.Info("rpc stream shutting down")
		// Let's not propagate this error further as we've already cancelled the context
		err = nil
	} else {
		log.Error(err, "rpc stream failed prematurely, will restart in background")
	}
	if err := stream.CloseSend(); err != nil {
		log.Error(err, "CloseSend returned an error")
	}
	if err != nil {
		errorCh <- err
	}
}

func (c *client) NeedLeaderElection() bool {
	return true
}

func tlsConfig(rootCaFile string, skipVerify bool) (*tls.Config, error) {
	// #nosec G402 -- we let the user decide if they want to ignore verification
	tlsConfig := &tls.Config{
		InsecureSkipVerify: skipVerify,
		MinVersion:         tls.VersionTLS12,
	}
	if rootCaFile != "" {
		roots := x509.NewCertPool()
		caCert, err := os.ReadFile(rootCaFile)
		if err != nil {
			return nil, errors.Wrapf(err, "could not read certificate %s", rootCaFile)
		}
		ok := roots.AppendCertsFromPEM(caCert)
		if !ok {
			return nil, errors.New("failed to parse root certificate")
		}
		tlsConfig.RootCAs = roots
	}
	return tlsConfig, nil
}
