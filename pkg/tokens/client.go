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
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"

	error_types "github.com/apache/dubbo-kubernetes/pkg/core/rest/errors/types"
	util_http "github.com/apache/dubbo-kubernetes/pkg/util/http"
)

type TokenClient struct {
	client util_http.Client
	url    string
}

func NewTokenClient(client util_http.Client, entity string) TokenClient {
	return TokenClient{
		client: client,
		url:    "/tokens/" + entity,
	}
}

func (tc TokenClient) Generate(tokenReq any) (string, error) {
	reqBytes, err := json.Marshal(tokenReq)
	if err != nil {
		return "", errors.Wrap(err, "could not marshal token request to json")
	}
	req, err := http.NewRequest(http.MethodPost, tc.url, bytes.NewReader(reqBytes))
	if err != nil {
		return "", errors.Wrap(err, "could not construct the request")
	}
	req.Header.Set("content-type", "application/json")
	resp, err := tc.client.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "could not execute the request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "could not read a body of the request")
	}
	if resp.StatusCode != http.StatusOK {
		var kumaErr error_types.Error
		if err := json.Unmarshal(body, &kumaErr); err == nil {
			if kumaErr.Title != "" {
				return "", &kumaErr
			}
		}
		return "", errors.Errorf("(%d): %s", resp.StatusCode, body)
	}
	return string(body), nil
}
