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
	"fmt"

	"github.com/rgeraldes24/go-qrl-beacon-client/spec/zond"
	bitfield "github.com/theQRL/go-bitfield"
)

// VersionedAttestation contains a versioned attestation.
type VersionedAttestation struct {
	Version        DataVersion
	ValidatorIndex *zond.ValidatorIndex
	Zond           *zond.Attestation
}

// IsEmpty returns true if there is no block.
func (v *VersionedAttestation) IsEmpty() bool {
	return v.Zond == nil
}

// AggregationBits returns the aggregation bits of the attestation.
func (v *VersionedAttestation) AggregationBits() (bitfield.Bitlist, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return nil, errors.New("no Zond attestation")
		}

		return v.Zond.AggregationBits, nil
	default:
		return nil, errors.New("unknown version")
	}
}

// Data returns the data of the attestation.
func (v *VersionedAttestation) Data() (*zond.AttestationData, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return nil, errors.New("no Zond attestation")
		}

		return v.Zond.Data, nil
	default:
		return nil, fmt.Errorf("unknown version: %d", v.Version)
	}
}

// CommitteeBits returns the committee bits of the attestation.
func (v *VersionedAttestation) CommitteeBits() (bitfield.Bitvector64, error) {
	switch v.Version {
	case DataVersionZond:
		return nil, errors.New("attestation does not provide committee bits")
	default:
		return nil, errors.New("unknown version")
	}
}

// CommitteeIndex returns the index if only one bit is set, otherwise error.
func (v *VersionedAttestation) CommitteeIndex() (zond.CommitteeIndex, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return 0, errors.New("no Zond attestation")
		}

		return v.Zond.Data.Index, nil
	default:
		return 0, errors.New("unknown version")
	}
}

func (v *VersionedAttestation) HashTreeRoot() ([32]byte, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return [32]byte{}, errors.New("no Zond attestation")
		}

		return v.Zond.HashTreeRoot()
	default:
		return [32]byte{}, errors.New("unknown version")
	}
}

// Signatures returns the signatures of the attestation.
func (v *VersionedAttestation) Signatures() ([]zond.MLDSA87Signature, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return nil, errors.New("no Zond attestation")
		}

		return v.Zond.Signatures, nil
	default:
		return nil, errors.New("unknown version")
	}
}

// String returns a string version of the structure.
func (v *VersionedAttestation) String() string {
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
