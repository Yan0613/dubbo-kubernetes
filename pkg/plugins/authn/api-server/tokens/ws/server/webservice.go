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

package server

import (
	"net/http"
	"time"

	"github.com/emicklei/go-restful/v3"

	"github.com/apache/dubbo-kubernetes/pkg/core"
	"github.com/apache/dubbo-kubernetes/pkg/core/rest/errors"
	"github.com/apache/dubbo-kubernetes/pkg/core/user"
	"github.com/apache/dubbo-kubernetes/pkg/core/validators"
	"github.com/apache/dubbo-kubernetes/pkg/plugins/authn/api-server/tokens/access"
	"github.com/apache/dubbo-kubernetes/pkg/plugins/authn/api-server/tokens/issuer"
	"github.com/apache/dubbo-kubernetes/pkg/plugins/authn/api-server/tokens/ws"
)

var log = core.Log.WithName("user-token-ws")

type userTokenWebService struct {
	issuer issuer.UserTokenIssuer
	access access.GenerateUserTokenAccess
}

func NewWebService(issuer issuer.UserTokenIssuer, access access.GenerateUserTokenAccess) *restful.WebService {
	webservice := userTokenWebService{
		issuer: issuer,
		access: access,
	}
	return webservice.createWs()
}

func (d *userTokenWebService) createWs() *restful.WebService {
	webservice := new(restful.WebService).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	webservice.Path("/tokens/user").
		Route(webservice.POST("").To(d.handleIdentityRequest))
	return webservice
}

func (d *userTokenWebService) handleIdentityRequest(request *restful.Request, response *restful.Response) {
	if err := d.access.ValidateGenerate(user.FromCtx(request.Request.Context())); err != nil {
		errors.HandleError(request.Request.Context(), response, err, "Could not issue a token")
		return
	}

	idReq := ws.UserTokenRequest{}
	if err := request.ReadEntity(&idReq); err != nil {
		log.Error(err, "Could not read a request")
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	verr := validators.ValidationError{}
	if idReq.Name == "" {
		verr.AddViolation("name", "cannot be empty")
	}

	var validFor time.Duration
	if idReq.ValidFor == "" {
		verr.AddViolation("validFor", "cannot be empty")
	} else {
		dur, err := time.ParseDuration(idReq.ValidFor)
		if err != nil {
			verr.AddViolation("validFor", "is invalid: "+err.Error())
		}
		validFor = dur
		if validFor == 0 {
			verr.AddViolation("validFor", "cannot be empty or nil")
		}
	}

	if verr.HasViolations() {
		errors.HandleError(request.Request.Context(), response, verr.OrNil(), "Invalid request")
		return
	}

	token, err := d.issuer.Generate(request.Request.Context(), user.User{
		Name:   idReq.Name,
		Groups: idReq.Groups,
	}, validFor)
	if err != nil {
		errors.HandleError(request.Request.Context(), response, err, "Could not issue a token")
		return
	}

	response.Header().Set("content-type", "text/plain")
	if _, err := response.Write([]byte(token)); err != nil {
		log.Error(err, "Could write a response")
	}
}
