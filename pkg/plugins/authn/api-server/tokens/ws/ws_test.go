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

package ws_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/emicklei/go-restful/v3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	store_config "github.com/apache/dubbo-kubernetes/pkg/config/core/resources/store"
	"github.com/apache/dubbo-kubernetes/pkg/core"
	"github.com/apache/dubbo-kubernetes/pkg/core/resources/manager"
	core_tokens "github.com/apache/dubbo-kubernetes/pkg/core/tokens"
	"github.com/apache/dubbo-kubernetes/pkg/core/user"
	"github.com/apache/dubbo-kubernetes/pkg/plugins/authn/api-server/tokens/issuer"
	"github.com/apache/dubbo-kubernetes/pkg/plugins/authn/api-server/tokens/ws/client"
	"github.com/apache/dubbo-kubernetes/pkg/plugins/authn/api-server/tokens/ws/server"
	"github.com/apache/dubbo-kubernetes/pkg/plugins/resources/memory"
	"github.com/apache/dubbo-kubernetes/pkg/test/matchers"
	util_http "github.com/apache/dubbo-kubernetes/pkg/util/http"
)

type noopGenerateUserTokenAccess struct{}

func (n *noopGenerateUserTokenAccess) ValidateGenerate(user.User) error {
	return nil
}

var _ = Describe("Auth Tokens WS", func() {
	var userTokenClient client.UserTokenClient
	var userTokenValidator issuer.UserTokenValidator
	var httpClient util_http.Client

	BeforeEach(func() {
		resManager := manager.NewResourceManager(memory.NewStore())
		signingKeyManager := core_tokens.NewSigningKeyManager(resManager, issuer.UserTokenSigningKeyPrefix)
		tokenIssuer := issuer.NewUserTokenIssuer(core_tokens.NewTokenIssuer(signingKeyManager))
		userTokenValidator = issuer.NewUserTokenValidator(
			core_tokens.NewValidator(
				core.Log.WithName("test"),
				[]core_tokens.SigningKeyAccessor{
					core_tokens.NewSigningKeyAccessor(resManager, issuer.UserTokenSigningKeyPrefix),
				},
				core_tokens.NewRevocations(resManager, issuer.UserTokenRevocationsGlobalSecretKey),
				store_config.MemoryStore,
			),
		)

		Expect(signingKeyManager.CreateDefaultSigningKey(context.Background())).To(Succeed())
		ws := server.NewWebService(tokenIssuer, &noopGenerateUserTokenAccess{})

		container := restful.NewContainer()
		container.Add(ws)
		srv := httptest.NewServer(container)

		baseURL, err := url.Parse(srv.URL)
		Expect(err).ToNot(HaveOccurred())
		httpClient = util_http.ClientWithBaseURL(http.DefaultClient, baseURL, nil)
		userTokenClient = client.NewHTTPUserTokenClient(httpClient)

		// wait for the server
		Eventually(func() error {
			_, err := userTokenClient.Generate("john.doe@example.com", []string{"team-a"}, time.Hour)
			return err
		}).ShouldNot(HaveOccurred())
	})

	It("should generate token", func() {
		// when
		token, err := userTokenClient.Generate("john.doe@example.com", []string{"team-a"}, 1*time.Hour)

		// then
		Expect(err).ToNot(HaveOccurred())
		u, err := userTokenValidator.Validate(context.Background(), token)
		Expect(err).ToNot(HaveOccurred())
		Expect(u.Name).To(Equal("john.doe@example.com"))
		Expect(u.Groups).To(Equal([]string{"team-a"}))
	})

	It("should throw an error when name is not passed", func() {
		// when
		_, err := userTokenClient.Generate("", nil, 1*time.Hour)

		// then
		bytes, err := json.MarshalIndent(err, "", "  ")
		Expect(err).ToNot(HaveOccurred())
		Expect(bytes).To(matchers.MatchGoldenJSON(path.Join("testdata", "ws-no-name.golden.json")))
	})

	It("should throw an error with 0 for validFor", func() {
		// when
		_, err := userTokenClient.Generate("foo@example.com", nil, 0)

		// then
		bytes, err := json.MarshalIndent(err, "", "  ")
		Expect(err).ToNot(HaveOccurred())
		Expect(bytes).To(matchers.MatchGoldenJSON(path.Join("testdata", "ws-0-validFor.golden.json")))
	})

	It("should throw an error if validFor is not present", func() {
		// given invalid request (cannot be implemented using UserTokenClient)
		req, err := http.NewRequest("POST", "/tokens/user", strings.NewReader(`{"name": "xyz"}`))
		req.Header.Add("content-type", "application/json")
		Expect(err).ToNot(HaveOccurred())

		// when
		resp, err := httpClient.Do(req)
		Expect(err).ToNot(HaveOccurred())
		defer resp.Body.Close()

		// then
		bytes, err := io.ReadAll(resp.Body)
		Expect(err).ToNot(HaveOccurred())
		Expect(bytes).To(matchers.MatchGoldenJSON(path.Join("testdata", "ws-missing-validFor.golden.json")))
	})

	It("should throw an error when issuer is disabled", func() {
		container := restful.NewContainer()
		ws := server.NewWebService(issuer.DisabledIssuer{}, &noopGenerateUserTokenAccess{})
		container.Add(ws)
		srv := httptest.NewServer(container)

		baseURL, err := url.Parse(srv.URL)
		Expect(err).ToNot(HaveOccurred())
		httpClient := util_http.ClientWithBaseURL(http.DefaultClient, baseURL, nil)
		userTokenClient := client.NewHTTPUserTokenClient(httpClient)

		Eventually(func(g Gomega) {
			_, err := userTokenClient.Generate("john.doe@example.com", []string{"team-a"}, 1*time.Hour)
			bytes, err := json.MarshalIndent(err, "", "  ")
			Expect(err).ToNot(HaveOccurred())
			Expect(bytes).To(matchers.MatchGoldenJSON(path.Join("testdata", "ws-token-issuer-disabled.golden.json")))
		}, "10s", "100ms").Should(Succeed())
	})
})
