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

package tokens

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/sethvargo/go-retry"
	"google.golang.org/protobuf/types/known/wrapperspb"

	dubbo_cp "github.com/apache/dubbo-kubernetes/pkg/config/app/dubbo-cp"
	config_core "github.com/apache/dubbo-kubernetes/pkg/config/core"
	"github.com/apache/dubbo-kubernetes/pkg/core/resources/apis/system"
	"github.com/apache/dubbo-kubernetes/pkg/core/resources/manager"
	"github.com/apache/dubbo-kubernetes/pkg/core/resources/model"
	core_store "github.com/apache/dubbo-kubernetes/pkg/core/resources/store"
	"github.com/apache/dubbo-kubernetes/pkg/core/runtime/component"
	"github.com/apache/dubbo-kubernetes/pkg/core/user"
	"github.com/apache/dubbo-kubernetes/pkg/plugins/authn/api-server/tokens/issuer"
)

var AdminTokenKey = model.ResourceKey{
	Name: "admin-user-token",
}

type adminTokenBootstrap struct {
	issuer     issuer.UserTokenIssuer
	resManager manager.ResourceManager
	cpCfg      dubbo_cp.Config
}

func NewAdminTokenBootstrap(issuer issuer.UserTokenIssuer, resManager manager.ResourceManager, cpCfg dubbo_cp.Config) component.Component {
	return &adminTokenBootstrap{
		issuer:     issuer,
		resManager: resManager,
		cpCfg:      cpCfg,
	}
}

func (a *adminTokenBootstrap) Start(stop <-chan struct{}) error {
	ctx, cancelFn := context.WithCancel(user.Ctx(context.Background(), user.ControlPlane))
	go func() {
		if err := a.generateTokenIfNotExist(ctx); err != nil {
			// just log, do not exist control plane
			log.Error(err, "could not generate Admin User Token")
		} else {
			msg := "bootstrap of Admin User Token is enabled. "
			if a.cpCfg.DeployMode == config_core.KubernetesMode {
				msg += fmt.Sprintf("To extract credentials execute 'kubectl get secret %s -n %s --template={{.data.value}} | base64 -d'. ", AdminTokenKey.Name, a.cpCfg.Store.Kubernetes.SystemNamespace)
			} else {
				msg += fmt.Sprintf("To extract admin credentials execute 'curl http://localhost:%d/global-secrets/%s | jq -r .data | base64 -d'. ", a.cpCfg.ApiServer.HTTP.Port, AdminTokenKey.Name)
			}
			msg += "You configure dubboctl with them 'dubboctl config control-planes add --auth-type=tokens --auth-conf token=YOUR_TOKEN'." +
				" To disable bootstrap of Admin User Token set dubbo_API_SERVER_AUTHN_TOKENS_BOOTSTRAP_ADMIN_TOKEN to false."
			log.Info(msg)
		}
	}()
	<-stop
	cancelFn()
	return nil
}

func (a *adminTokenBootstrap) generateTokenIfNotExist(ctx context.Context) error {
	secret := system.NewGlobalSecretResource()
	err := a.resManager.Get(ctx, secret, core_store.GetBy(AdminTokenKey))
	if err == nil {
		return nil // already exists
	}
	if !core_store.IsResourceNotFound(err) {
		return errors.Wrap(err, "could not check if token exist")
	}

	log.Info("Admin User Token not found. Generating.")
	token, err := a.generateAdminToken(ctx)
	if err != nil {
		return errors.Wrap(err, "could not generate admin token")
	}

	log.Info("saving generated Admin User Token", "globalSecretName", AdminTokenKey.Name)
	secret.Spec.Data = &wrapperspb.BytesValue{
		Value: []byte(token),
	}
	if err := a.resManager.Create(ctx, secret, core_store.CreateBy(AdminTokenKey)); err != nil {
		return err
	}
	return nil
}

func (a *adminTokenBootstrap) generateAdminToken(ctx context.Context) (string, error) {
	// we need retries because signing key may not be available yet
	var token string
	err := retry.Do(ctx, retry.WithMaxDuration(10*time.Minute, retry.NewConstant(time.Second)), func(ctx context.Context) error {
		t, err := a.issuer.Generate(ctx, user.Admin, 24*365*10*time.Hour)
		if err != nil {
			return retry.RetryableError(err)
		}
		token = t
		return nil
	})
	return token, err
}

func (a *adminTokenBootstrap) NeedLeaderElection() bool {
	return true
}

var _ component.Component = &adminTokenBootstrap{}
