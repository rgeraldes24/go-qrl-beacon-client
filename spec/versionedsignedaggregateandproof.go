// Copyright © 2025 Attestant Limited.
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

// VersionedSignedAggregateAndProof contains a versioned signed aggregate and proof.
type VersionedSignedAggregateAndProof struct {
	Version DataVersion
	Zond    *zond.SignedAggregateAndProof
}

// AggregatorIndex returns the aggregator index of the aggregate.
func (v *VersionedSignedAggregateAndProof) AggregatorIndex() (zond.ValidatorIndex, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return 0, errors.New("no zond signed aggregate and proof")
		}

		return v.Zond.Message.AggregatorIndex, nil
	default:
		return 0, errors.New("unknown version for signed aggregate and proof")
	}
}

// IsEmpty returns true if there is no aggregate and proof.
func (v *VersionedSignedAggregateAndProof) IsEmpty() bool {
	return v.Zond == nil
}

// SelectionProof returns the selection proof of the signed aggregate.
func (v *VersionedSignedAggregateAndProof) SelectionProof() (zond.MLDSA87Signature, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return zond.MLDSA87Signature{}, errors.New("no zond signed aggregate and proof")
		}

		return v.Zond.Message.SelectionProof, nil
	default:
		return zond.MLDSA87Signature{}, errors.New("unknown version")
	}
}

// Signature returns the signature of the signed aggregate and proof.
func (v *VersionedSignedAggregateAndProof) Signature() (zond.MLDSA87Signature, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return zond.MLDSA87Signature{}, errors.New("no zond signed aggregate and proof")
		}

		return v.Zond.Signature, nil
	default:
		return zond.MLDSA87Signature{}, errors.New("unknown version")
	}
}

// Slot returns the slot of the signed aggregate and proof.
func (v *VersionedSignedAggregateAndProof) Slot() (zond.Slot, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return 0, errors.New("no zond signed aggregate and proof")
		}

		return v.Zond.Message.Aggregate.Data.Slot, nil
	default:
		return 0, errors.New("unknown version")
	}
}

// String returns a string version of the structure.
func (v *VersionedSignedAggregateAndProof) String() string {
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
