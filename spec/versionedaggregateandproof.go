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

// VersionedAggregateAndProof contains a versioned aggregate and proof.
type VersionedAggregateAndProof struct {
	Version DataVersion
	Zond    *zond.AggregateAndProof
}

// AggregatorIndex returns the aggregator index of the aggregate.
func (v *VersionedAggregateAndProof) AggregatorIndex() (zond.ValidatorIndex, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return 0, errors.New("no zond aggregate and proof")
		}

		return v.Zond.AggregatorIndex, nil
	default:
		return 0, errors.New("unknown version for aggregate and proof")
	}
}

// HashTreeRoot returns the hash tree root of the aggregate and proof.
func (v *VersionedAggregateAndProof) HashTreeRoot() ([32]byte, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return [32]byte{}, errors.New("no zond aggregate and proof")
		}

		return v.Zond.HashTreeRoot()
	default:
		return [32]byte{}, errors.New("unknown version")
	}
}

// IsEmpty returns true if there is no aggregate and proof.
func (v *VersionedAggregateAndProof) IsEmpty() bool {
	return v.Zond == nil
}

// String returns a string version of the structure.
func (v *VersionedAggregateAndProof) String() string {
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

// SelectionProof returns the selection proof of the aggregate.
func (v *VersionedAggregateAndProof) SelectionProof() (zond.MLDSA87Signature, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return zond.MLDSA87Signature{}, errors.New("no zond aggregate and proof")
		}

		return v.Zond.SelectionProof, nil
	default:
		return zond.MLDSA87Signature{}, errors.New("unknown version")
	}
}
