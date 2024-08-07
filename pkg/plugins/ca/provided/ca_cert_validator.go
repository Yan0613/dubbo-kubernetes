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

package provided

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"

	"github.com/apache/dubbo-kubernetes/pkg/core/validators"
	util_tls "github.com/apache/dubbo-kubernetes/pkg/tls"
)

func ValidateCaCert(signingPair util_tls.KeyPair) error {
	err := validateCaCert(signingPair)
	return err.OrNil()
}

func validateCaCert(signingPair util_tls.KeyPair) validators.ValidationError {
	var verr validators.ValidationError
	tlsKeyPair, err := tls.X509KeyPair(signingPair.CertPEM, signingPair.KeyPEM)
	if err != nil {
		verr.AddViolation("cert", fmt.Sprintf("not a valid TLS key pair: %s", err))
		return verr
	}
	for i, certificate := range tlsKeyPair.Certificate {
		path := validators.RootedAt("cert").Index(i)
		cert, err := x509.ParseCertificate(certificate)
		if err != nil {
			verr.AddViolationAt(path, fmt.Sprintf("not a valid x509 certificate: %s", err))
			return verr
		}
		if !cert.IsCA {
			verr.AddViolationAt(path, "basic constraint 'CA' must be set to 'true' (see X509-SVID: 4.1. Basic Constraints)")
		}
		if cert.KeyUsage&x509.KeyUsageCertSign == 0 {
			verr.AddViolationAt(path, "key usage extension 'keyCertSign' must be set (see X509-SVID: 4.3. Key Usage)")
		}
		if cert.KeyUsage&x509.KeyUsageKeyAgreement != 0 {
			verr.AddViolationAt(path, "key usage extension 'keyAgreement' must NOT be set (see X509-SVID: Appendix A. X.509 Field Reference)")
		}
	}
	return verr
}
