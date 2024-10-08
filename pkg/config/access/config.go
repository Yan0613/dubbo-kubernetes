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

package access

import (
	"github.com/pkg/errors"

	"github.com/apache/dubbo-kubernetes/pkg/config"
)

const StaticType = "static"

func DefaultAccessConfig() AccessConfig {
	return AccessConfig{
		Type: StaticType,
		Static: StaticAccessConfig{
			AdminResources: AdminResourcesStaticAccessConfig{
				Users:  []string{"mesh-system:admin"},
				Groups: []string{"mesh-system:admin"},
			},
			GenerateDPToken: GenerateDPTokenStaticAccessConfig{
				Users:  []string{"mesh-system:admin"},
				Groups: []string{"mesh-system:admin"},
			},
			GenerateUserToken: GenerateUserTokenStaticAccessConfig{
				Users:  []string{"mesh-system:admin"},
				Groups: []string{"mesh-system:admin"},
			},
			GenerateZoneToken: GenerateZoneTokenStaticAccessConfig{
				Users:  []string{"mesh-system:admin"},
				Groups: []string{"mesh-system:admin"},
			},
			ViewConfigDump: ViewConfigDumpStaticAccessConfig{
				Users:  []string{},
				Groups: []string{"mesh-system:unauthenticated", "mesh-system:authenticated"},
			},
			ViewStats: ViewStatsStaticAccessConfig{
				Users:  []string{},
				Groups: []string{"mesh-system:unauthenticated", "mesh-system:authenticated"},
			},
			ViewClusters: ViewClustersStaticAccessConfig{
				Users:  []string{},
				Groups: []string{"mesh-system:unauthenticated", "mesh-system:authenticated"},
			},
			ControlPlaneMetadata: ControlPlaneMetadataStaticAccessConfig{
				Users:  []string{},
				Groups: []string{"mesh-system:unauthenticated", "mesh-system:authenticated"},
			},
		},
	}
}

// AccessConfig defines a configuration for access control
type AccessConfig struct {
	config.BaseConfig

	// Type of the access strategy (available values: "static")
	Type string `json:"type" envconfig:"DUBBO_ACCESS_TYPE"`
	// Configuration of static access strategy
	Static StaticAccessConfig `json:"static"`
}

func (r AccessConfig) Validate() error {
	if r.Type == "" {
		return errors.New("Type has to be defined")
	}
	return nil
}

var _ config.Config = &AccessConfig{}

// StaticAccessConfig a static access strategy configuration
type StaticAccessConfig struct {
	// AdminResources defines an access to admin resources (Secret/GlobalSecret)
	AdminResources AdminResourcesStaticAccessConfig `json:"adminResources"`
	// GenerateDPToken defines an access to generating dataplane token
	GenerateDPToken GenerateDPTokenStaticAccessConfig `json:"generateDpToken"`
	// GenerateUserToken defines an access to generating user token
	GenerateUserToken GenerateUserTokenStaticAccessConfig `json:"generateUserToken"`
	// GenerateZoneToken defines an access to generating zone token
	GenerateZoneToken GenerateZoneTokenStaticAccessConfig `json:"generateZoneToken"`
	// ViewConfigDump defines an access to getting envoy config dump
	ViewConfigDump ViewConfigDumpStaticAccessConfig `json:"viewConfigDump"`
	// ViewStats defines an access to getting envoy stats
	ViewStats ViewStatsStaticAccessConfig `json:"viewStats"`
	// ViewClusters defines an access to getting envoy clusters
	ViewClusters ViewClustersStaticAccessConfig `json:"viewClusters"`
	// ControlPlaneMetadata defines an access for control-plane metadata (for example config)
	ControlPlaneMetadata ControlPlaneMetadataStaticAccessConfig `json:"controlPlaneMetadata"`
}

type AdminResourcesStaticAccessConfig struct {
	// List of users that are allowed to access admin resources
	Users []string `json:"users" envconfig:"DUBBO_ACCESS_STATIC_ADMIN_RESOURCES_USERS"`
	// List of groups that are allowed to access admin resources
	Groups []string `json:"groups" envconfig:"DUBBO_ACCESS_STATIC_ADMIN_RESOURCES_GROUPS"`
}

type GenerateDPTokenStaticAccessConfig struct {
	// List of users that are allowed to generate dataplane token
	Users []string `json:"users" envconfig:"DUBBO_ACCESS_STATIC_GENERATE_DP_TOKEN_USERS"`
	// List of groups that are allowed to generate dataplane token
	Groups []string `json:"groups" envconfig:"DUBBO_ACCESS_STATIC_GENERATE_DP_TOKEN_GROUPS"`
}

type GenerateUserTokenStaticAccessConfig struct {
	// List of users that are allowed to generate user token
	Users []string `json:"users" envconfig:"DUBBO_ACCESS_STATIC_GENERATE_USER_TOKEN_USERS"`
	// List of groups that are allowed to generate user token
	Groups []string `json:"groups" envconfig:"DUBBO_ACCESS_STATIC_GENERATE_USER_TOKEN_GROUPS"`
}

type GenerateZoneTokenStaticAccessConfig struct {
	// List of users that are allowed to generate zone token
	Users []string `json:"users" envconfig:"DUBBO_ACCESS_STATIC_GENERATE_ZONE_TOKEN_USERS"`
	// List of groups that are allowed to generate zone token
	Groups []string `json:"groups" envconfig:"DUBBO_ACCESS_STATIC_GENERATE_ZONE_TOKEN_GROUPS"`
}

type ViewConfigDumpStaticAccessConfig struct {
	// List of users that are allowed to get envoy config dump
	Users []string `json:"users" envconfig:"DUBBO_ACCESS_STATIC_GET_CONFIG_DUMP_USERS"`
	// List of groups that are allowed to get envoy config dump
	Groups []string `json:"groups" envconfig:"DUBBO_ACCESS_STATIC_GET_CONFIG_DUMP_GROUPS"`
}

type ViewStatsStaticAccessConfig struct {
	// List of users that are allowed to get envoy config stats
	Users []string `json:"users" envconfig:"DUBBO_ACCESS_STATIC_VIEW_STATS_USERS"`
	// List of groups that are allowed to get envoy config stats
	Groups []string `json:"groups" envconfig:"DUBBO_ACCESS_STATIC_VIEW_STATS_GROUPS"`
}

type ViewClustersStaticAccessConfig struct {
	// List of users that are allowed to get envoy config clusters
	Users []string `json:"users" envconfig:"DUBBO_ACCESS_STATIC_VIEW_CLUSTERS_USERS"`
	// List of groups that are allowed to get envoy config clusters
	Groups []string `json:"groups" envconfig:"DUBBO_ACCESS_STATIC_VIEW_CLUSTERS_GROUPS"`
}

type ControlPlaneMetadataStaticAccessConfig struct {
	// List of users that are allowed to access control-plane metadata
	Users []string `json:"users" envconfig:"DUBBO_ACCESS_STATIC_CONTROL_PLANE_METADATA_USERS"`
	// List of groups that are allowed to access control-plane metadata
	Groups []string `json:"groups" envconfig:"DUBBO_ACCESS_STATIC_CONTROL_PLANE_METADATA_GROUPS"`
}
