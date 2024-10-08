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
	"github.com/apache/dubbo-kubernetes/pkg/core"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"time"
)

// Issuer generates tokens.
// Token is a JWT token with claims that is provided by the actual issuer (for example - Dataplane Token Issuer, User Token Issuer).
// We place "kid" in token, so we don't have to validate the token against every single signing key.
// Instead, we take "kid" from the token, retrieve signing key and validate only against this key.
// A new token is always generated by using the latest signing key.
type Issuer interface {
	Generate(ctx context.Context, claims Claims, validFor time.Duration) (Token, error)
}

type jwtTokenIssuer struct {
	signingKeyManager SigningKeyManager
}

func NewTokenIssuer(signingKeyAccessor SigningKeyManager) Issuer {
	return &jwtTokenIssuer{
		signingKeyManager: signingKeyAccessor,
	}
}

var _ Issuer = &jwtTokenIssuer{}

func (j *jwtTokenIssuer) Generate(ctx context.Context, claims Claims, validFor time.Duration) (Token, error) {
	signingKey, keyID, err := j.signingKeyManager.GetLatestSigningKey(ctx)
	if err != nil {
		return "", err
	}

	now := core.Now()
	claims.SetRegisteredClaims(jwt.RegisteredClaims{
		ID:        core.NewUUID(),
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now.Add(time.Minute * -5)), // todo(jakubdyszkiewicz) parametrize via config and go through all clock skews in the project
		ExpiresAt: jwt.NewNumericDate(now.Add(validFor)),
	})

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header[KeyIDHeader] = keyID
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", errors.Wrap(err, "could not sign a token")
	}
	return tokenString, nil
}

var IssuerDisabled = errors.New("issuing tokens using the control plane is disabled.")
