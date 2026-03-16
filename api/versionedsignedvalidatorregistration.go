// Copyright © 2022 Attestant Limited.
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

package api

import (
	"time"

	apiv1 "github.com/theQRL/go-qrl-beacon-client/api/v1"
	"github.com/theQRL/go-qrl-beacon-client/spec"
	"github.com/theQRL/go-qrl-beacon-client/spec/zond"
)

// VersionedSignedValidatorRegistration contains a versioned SignedValidatorRegistrationV1.
type VersionedSignedValidatorRegistration struct {
	Version spec.BuilderVersion                `json:"version"`
	V1      *apiv1.SignedValidatorRegistration `json:"v1"`
}

// FeeRecipient returns the fee recipient of the signed validator registration.
func (v *VersionedSignedValidatorRegistration) FeeRecipient() (zond.ExecutionAddress, error) {
	switch v.Version {
	case spec.BuilderVersionV1:
		if v.V1 == nil {
			return zond.ExecutionAddress{}, ErrDataMissing
		}

		return v.V1.Message.FeeRecipient, nil
	default:
		return zond.ExecutionAddress{}, ErrUnsupportedVersion
	}
}

// GasLimit returns the gas limit of the signed validator registration.
func (v *VersionedSignedValidatorRegistration) GasLimit() (uint64, error) {
	switch v.Version {
	case spec.BuilderVersionV1:
		if v.V1 == nil {
			return 0, ErrDataMissing
		}

		return v.V1.Message.GasLimit, nil
	default:
		return 0, ErrUnsupportedVersion
	}
}

// Timestamp returns the timestamp of the signed validator registration.
func (v *VersionedSignedValidatorRegistration) Timestamp() (time.Time, error) {
	switch v.Version {
	case spec.BuilderVersionV1:
		if v.V1 == nil {
			return time.Time{}, ErrDataMissing
		}

		return v.V1.Message.Timestamp, nil
	default:
		return time.Time{}, ErrUnsupportedVersion
	}
}

// PubKey returns the public key of the signed validator registration.
func (v *VersionedSignedValidatorRegistration) PubKey() (zond.MLDSA87PubKey, error) {
	switch v.Version {
	case spec.BuilderVersionV1:
		if v.V1 == nil {
			return zond.MLDSA87PubKey{}, ErrDataMissing
		}

		return v.V1.Message.Pubkey, nil
	default:
		return zond.MLDSA87PubKey{}, ErrUnsupportedVersion
	}
}

// Root returns the root of the validator registration.
func (v *VersionedSignedValidatorRegistration) Root() (zond.Root, error) {
	switch v.Version {
	case spec.BuilderVersionV1:
		if v.V1 == nil {
			return zond.Root{}, ErrDataMissing
		}

		return v.V1.Message.HashTreeRoot()
	default:
		return zond.Root{}, ErrUnsupportedVersion
	}
}
