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
	envoy_upstream_http "github.com/envoyproxy/go-control-plane/envoy/extensions/upstreams/http/v3"

	mesh_proto "github.com/apache/dubbo-kubernetes/api/mesh/v1alpha1"
	core_mesh "github.com/apache/dubbo-kubernetes/pkg/core/resources/apis/mesh"
	policies_defaults "github.com/apache/dubbo-kubernetes/pkg/plugins/policies/core/defaults"
	util_proto "github.com/apache/dubbo-kubernetes/pkg/util/proto"
)

type TimeoutConfigurer struct {
	Protocol core_mesh.Protocol
	Conf     *mesh_proto.Timeout_Conf
}

var _ ClusterConfigurer = &TimeoutConfigurer{}

func (t *TimeoutConfigurer) Configure(cluster *envoy_cluster.Cluster) error {
	cluster.ConnectTimeout = util_proto.Duration(t.Conf.GetConnectTimeoutOrDefault(policies_defaults.DefaultConnectTimeout))
	switch t.Protocol {
	case core_mesh.ProtocolHTTP, core_mesh.ProtocolHTTP2, core_mesh.ProtocolGRPC:
		err := UpdateCommonHttpProtocolOptions(cluster, func(options *envoy_upstream_http.HttpProtocolOptions) {
			if options.CommonHttpProtocolOptions == nil {
				options.CommonHttpProtocolOptions = &envoy_core.HttpProtocolOptions{}
			}

			t.setIdleTimeout(options.CommonHttpProtocolOptions)
			t.setMaxStreamDuration(options.CommonHttpProtocolOptions)
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *TimeoutConfigurer) setIdleTimeout(options *envoy_core.HttpProtocolOptions) {
	options.IdleTimeout = util_proto.Duration(t.Conf.GetHttp().GetIdleTimeout().AsDuration())
}

func (t *TimeoutConfigurer) setMaxStreamDuration(options *envoy_core.HttpProtocolOptions) {
	if msd := t.Conf.GetHttp().GetMaxStreamDuration(); msd != nil && msd.AsDuration() != 0 {
		options.MaxStreamDuration = msd
		return
	}

	// backwards compatibility
	if t.Protocol == core_mesh.ProtocolGRPC {
		if msd := t.Conf.GetGrpc().GetMaxStreamDuration(); msd != nil && msd.AsDuration() != 0 {
			options.MaxStreamDuration = util_proto.Duration(msd.AsDuration())
		}
	}
}
