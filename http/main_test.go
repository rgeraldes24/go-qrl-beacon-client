// Copyright © 2020 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package http_test

import (
	"encoding/hex"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/theQRL/go-qrl-beacon-client/spec/zond"
)

// timeout for tests.
// var timeout = 60 * time.Second
var timeout = 10 * time.Minute

func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.Disabled)
	if os.Getenv("HTTP_ADDRESS") != "" {
		os.Exit(m.Run())
	}
}

// mustParseRoot is used for testing.
func mustParseRoot(input string) *zond.Root {
	root, err := hex.DecodeString(strings.TrimPrefix(input, "0x"))
	if err != nil {
		panic("invalid root")
	}
	if len(root) != zond.RootLength {
		panic("invalid length root")
	}

	var res zond.Root
	copy(res[:], root)

	return &res
}

// mustParseSignature is used for testing.
func mustParseSignature(input string) *zond.MLDSA87Signature {
	sig, err := hex.DecodeString(strings.TrimPrefix(input, "0x"))
	if err != nil {
		panic("invalid signature")
	}
	if len(sig) != zond.SignatureLength {
		panic("invalid length signature")
	}

	var res zond.MLDSA87Signature
	copy(res[:], sig)

	return &res
}

// mustParsePubKey is used for testing.
func mustParsePubKey(input string) *zond.MLDSA87PubKey {
	pubKey, err := hex.DecodeString(strings.TrimPrefix(input, "0x"))
	if err != nil {
		panic("invalid public key")
	}
	if len(pubKey) != zond.PublicKeyLength {
		panic("invalid length public key")
	}

	var res zond.MLDSA87PubKey
	copy(res[:], pubKey)

	return &res
}
