// Copyright © 2024 Attestant Limited.
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

package spec

import (
	"errors"

	"github.com/theQRL/go-qrl-beacon-client/spec/zond"
)

// VersionedIndexedAttestation contains a versioned indexed attestation.
type VersionedIndexedAttestation struct {
	Version DataVersion
	Zond    *zond.IndexedAttestation
}

// IsEmpty returns true if there is no block.
func (v *VersionedIndexedAttestation) IsEmpty() bool {
	return v.Zond == nil
}

// AttestingIndices returns the attesting indices of the indexed attestation.
func (v *VersionedIndexedAttestation) AttestingIndices() ([]uint64, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return nil, errors.New("no Zond indexed attestation")
		}

		return v.Zond.AttestingIndices, nil
	default:
		return nil, errors.New("unknown version")
	}
}

// Data returns the data of the indexed attestation.
func (v *VersionedIndexedAttestation) Data() (*zond.AttestationData, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return nil, errors.New("no Zond indexed attestation")
		}

		return v.Zond.Data, nil
	default:
		return nil, errors.New("unknown version")
	}
}

// Signatures returns the signatures of the indexed attestation.
func (v *VersionedIndexedAttestation) Signatures() ([]zond.MLDSA87Signature, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return nil, errors.New("no Zond indexed attestation")
		}

		return v.Zond.Signatures, nil
	default:
		return nil, errors.New("unknown version")
	}
}

// String returns a string version of the structure.
func (v *VersionedIndexedAttestation) String() string {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return ""
		}

		return v.Zond.String()
	default:
		return "unknown version"
	}
}
