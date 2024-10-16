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

package clusters

import (
	envoy_cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	envoy_core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"google.golang.org/protobuf/types/known/structpb"

	mesh_proto "github.com/apache/dubbo-kubernetes/api/mesh/v1alpha1"
	core_mesh "github.com/apache/dubbo-kubernetes/pkg/core/resources/apis/mesh"
	core_xds "github.com/apache/dubbo-kubernetes/pkg/core/xds"
	"github.com/apache/dubbo-kubernetes/pkg/util/proto"
	envoy_metadata "github.com/apache/dubbo-kubernetes/pkg/xds/envoy/metadata/v3"
	"github.com/apache/dubbo-kubernetes/pkg/xds/envoy/tags"
	"github.com/apache/dubbo-kubernetes/pkg/xds/envoy/tls"
	envoy_tls "github.com/apache/dubbo-kubernetes/pkg/xds/envoy/tls/v3"
)

type ClientSideMTLSConfigurer struct {
	SecretsTracker   core_xds.SecretsTracker
	UpstreamMesh     *core_mesh.MeshResource
	UpstreamService  string
	LocalMesh        *core_mesh.MeshResource
	Tags             []tags.Tags
	SNI              string
	UpstreamTLSReady bool
	VerifyIdentities []string
}

var _ ClusterConfigurer = &ClientSideMTLSConfigurer{}

func (c *ClientSideMTLSConfigurer) Configure(cluster *envoy_cluster.Cluster) error {
	if !c.UpstreamMesh.MTLSEnabled() || !c.LocalMesh.MTLSEnabled() {
		return nil
	}
	if c.UpstreamMesh.GetEnabledCertificateAuthorityBackend().Mode == mesh_proto.CertificateAuthorityBackend_PERMISSIVE &&
		!c.UpstreamTLSReady {
		return nil
	}

	meshName := c.UpstreamMesh.GetMeta().GetName()
	// there might be a situation when there are multiple sam tags passed here for example two outbound listeners with the same tags, therefore we need to distinguish between them.
	distinctTags := tags.DistinctTags(c.Tags)
	switch {
	case len(distinctTags) == 0:
		transportSocket, err := c.createTransportSocket(c.SNI)
		if err != nil {
			return err
		}
		cluster.TransportSocket = transportSocket
	case len(distinctTags) == 1:
		sni := tls.SNIFromTags(c.Tags[0].WithTags("mesh", meshName))
		transportSocket, err := c.createTransportSocket(sni)
		if err != nil {
			return err
		}
		cluster.TransportSocket = transportSocket
	default:
		for _, tags := range distinctTags {
			sni := tls.SNIFromTags(tags.WithTags("mesh", meshName))
			transportSocket, err := c.createTransportSocket(sni)
			if err != nil {
				return err
			}
			cluster.TransportSocketMatches = append(cluster.TransportSocketMatches, &envoy_cluster.Cluster_TransportSocketMatch{
				Name: sni,
				Match: &structpb.Struct{
					Fields: envoy_metadata.MetadataFields(tags.WithoutTags(mesh_proto.ServiceTag)),
				},
				TransportSocket: transportSocket,
			})
		}
	}
	return nil
}

func (c *ClientSideMTLSConfigurer) createTransportSocket(sni string) (*envoy_core.TransportSocket, error) {
	if !c.UpstreamMesh.MTLSEnabled() {
		return nil, nil
	}

	ca := c.SecretsTracker.RequestCa(c.UpstreamMesh.GetMeta().GetName())
	identity := c.SecretsTracker.RequestIdentityCert()

	var verifyIdentities []string
	if c.VerifyIdentities != nil {
		verifyIdentities = c.VerifyIdentities
	}
	tlsContext, err := envoy_tls.CreateUpstreamTlsContext(identity, ca, c.UpstreamService, sni, verifyIdentities)
	if err != nil {
		return nil, err
	}
	if tlsContext == nil {
		return nil, nil
	}
	pbst, err := proto.MarshalAnyDeterministic(tlsContext)
	if err != nil {
		return nil, err
	}
	transportSocket := &envoy_core.TransportSocket{
		Name: "envoy.transport_sockets.tls",
		ConfigType: &envoy_core.TransportSocket_TypedConfig{
			TypedConfig: pbst,
		},
	}
	return transportSocket, nil
}
