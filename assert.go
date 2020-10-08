/*
 * Copyright 2019 Aletheia Ware LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package testinggo

import (
	"bytes"
	"crypto/rsa"
	"encoding/base64"
	"github.com/AletheiaWareLLC/cryptogo"
	"github.com/golang/protobuf/proto"
	"regexp"
	"testing"
)

func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Expected no error, instead got '%s'", err)
	}
}

func AssertError(t *testing.T, expected string, err error) {
	t.Helper()
	if err == nil {
		t.Fatalf("Expected error '%s'", expected)
	} else if err.Error() != expected {
		t.Fatalf("Expected error '%s', instead got '%s'", expected, err)
	}
}

func AssertMatchesError(t *testing.T, expected string, err error) {
	t.Helper()
	if err == nil {
		t.Fatalf("Expected error '%s'", expected)
	} else {
		match, _ := regexp.MatchString(expected, err.Error())
		if !match {
			t.Fatalf("Expected error '%s', instead got '%s'", expected, err)
		}
	}
}

func AssertHashEqual(t *testing.T, expected, actual []byte) {
	t.Helper()
	if !bytes.Equal(expected, actual) {
		t.Fatalf("Wrong hash: expected '%s' , instead got '%s'", base64.RawURLEncoding.EncodeToString(expected), base64.RawURLEncoding.EncodeToString(actual))
	}
}

func AssertProtobufEqual(t *testing.T, expected, actual proto.Message) {
	t.Helper()
	es := expected.String()
	as := actual.String()
	if es != as {
		t.Fatalf("Wrong protobuf: expected '%s' , instead got '%s'", es, as)
	}
}

func AssertPrivateKeyEqual(t *testing.T, expected, actual *rsa.PrivateKey) {
	t.Helper()
	hash := cryptogo.Hash([]byte("abcdefghijklmnopqrstuvwxyz"))
	expectedSignature, err := cryptogo.CreateSignature(expected, hash, cryptogo.SignatureAlgorithm_SHA512WITHRSA_PSS)
	AssertNoError(t, err)
	AssertNoError(t, cryptogo.VerifySignature(&actual.PublicKey, hash, expectedSignature, cryptogo.SignatureAlgorithm_SHA512WITHRSA_PSS))
	actualSignature, err := cryptogo.CreateSignature(actual, hash, cryptogo.SignatureAlgorithm_SHA512WITHRSA_PSS)
	AssertNoError(t, err)
	AssertNoError(t, cryptogo.VerifySignature(&expected.PublicKey, hash, actualSignature, cryptogo.SignatureAlgorithm_SHA512WITHRSA_PSS))
}
