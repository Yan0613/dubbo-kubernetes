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

package k8s

import (
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	secret_model "github.com/apache/dubbo-kubernetes/pkg/core/resources/apis/system"
	"github.com/apache/dubbo-kubernetes/pkg/core/resources/model"
	common_k8s "github.com/apache/dubbo-kubernetes/pkg/plugins/common/k8s"
	"github.com/apache/dubbo-kubernetes/pkg/plugins/resources/k8s"
	k8s_model "github.com/apache/dubbo-kubernetes/pkg/plugins/resources/k8s/native/pkg/model"
)

// NewKubernetesMapper creates a ResourceMapper that returns the k8s object as is. This is meant to be used when the underlying store is kubernetes
func NewKubernetesMapper() k8s.ResourceMapperFunc {
	return func(resource model.Resource, namespace string) (k8s_model.KubernetesObject, error) {
		res, err := DefaultConverter().ToKubernetesObject(resource)
		res.TypeMeta = metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		}
		if err != nil {
			return nil, err
		}
		if namespace != "" {
			res.SetNamespace(namespace)
		}
		return &Secret{Secret: *res}, nil
	}
}

// NewInferenceMapper creates a ResourceMapper that infers a k8s resource from the core_model. Extract namespace from the name if necessary.
// This mostly useful when the underlying store is not kubernetes but you want to show what a kubernetes version of the policy would be like (in global for example).
func NewInferenceMapper(systemNamespace string) k8s.ResourceMapperFunc {
	return func(resource model.Resource, namespace string) (k8s_model.KubernetesObject, error) {
		var rs k8s_model.KubernetesObject
		switch resource.Descriptor().Name {
		case secret_model.SecretType:
			rs = NewSecret(common_k8s.MeshSecretType)
			rs.SetMesh(resource.GetMeta().GetMesh())
		case secret_model.GlobalSecretType:
			rs = NewSecret(common_k8s.GlobalSecretType)
		default:
			return nil, errors.New("invalid resource type")
		}
		if namespace != "" { // If the user is forcing the namespace accept it.
			rs.SetNamespace(namespace)
		} else {
			rs.SetNamespace(systemNamespace)
		}
		rs.SetName(resource.GetMeta().GetName())
		rs.SetCreationTimestamp(v1.NewTime(resource.GetMeta().GetCreationTime()))
		rs.SetSpec(resource.GetSpec())
		return rs, nil
	}
}
