// Copyright © 2022, 2023 Attestant Limited.
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
	apiv1zond "github.com/theQRL/go-qrl-beacon-client/api/v1/zond"
	"github.com/theQRL/go-qrl-beacon-client/spec"
	"github.com/theQRL/go-qrl-beacon-client/spec/zond"
)

// VersionedBlindedProposal contains a versioned blinded proposal.
type VersionedBlindedProposal struct {
	Version spec.DataVersion
	Zond    *apiv1zond.BlindedBeaconBlock
}

// IsEmpty returns true if there is no proposal.
func (v *VersionedBlindedProposal) IsEmpty() bool {
	return v.Zond == nil
}

// Slot returns the slot of the blinded proposal.
func (v *VersionedBlindedProposal) Slot() (zond.Slot, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil {
			return 0, ErrDataMissing
		}

		return v.Zond.Slot, nil
	default:
		return 0, ErrUnsupportedVersion
	}
}

// ProposerIndex returns the proposer index of the blinded proposal.
func (v *VersionedBlindedProposal) ProposerIndex() (zond.ValidatorIndex, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil {
			return 0, ErrDataMissing
		}

		return v.Zond.ProposerIndex, nil
	default:
		return 0, ErrUnsupportedVersion
	}
}

// RandaoReveal returns the RANDAO reveal of the blinded proposal.
func (v *VersionedBlindedProposal) RandaoReveal() (zond.MLDSA87Signature, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Body == nil {
			return zond.MLDSA87Signature{}, ErrDataMissing
		}

		return v.Zond.Body.RANDAOReveal, nil
	default:
		return zond.MLDSA87Signature{}, ErrUnsupportedVersion
	}
}

// Graffiti returns the graffiti of the blinded proposal.
func (v *VersionedBlindedProposal) Graffiti() ([32]byte, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Body == nil {
			return [32]byte{}, ErrDataMissing
		}

		return v.Zond.Body.Graffiti, nil
	default:
		return [32]byte{}, ErrUnsupportedVersion
	}
}

// Attestations returns the attestations of the blinded proposal.
func (v *VersionedBlindedProposal) Attestations() ([]spec.VersionedAttestation, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil || v.Zond.Body == nil {
			return nil, ErrDataMissing
		}

		versionedAttestations := make([]spec.VersionedAttestation, len(v.Zond.Body.Attestations))
		for i, attestation := range v.Zond.Body.Attestations {
			versionedAttestations[i] = spec.VersionedAttestation{
				Version: spec.DataVersionZond,
				Zond:    attestation,
			}
		}

		return versionedAttestations, nil
	default:
		return nil, ErrUnsupportedVersion
	}
}

// Root returns the root of the blinded proposal.
func (v *VersionedBlindedProposal) Root() (zond.Root, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil {
			return zond.Root{}, ErrDataMissing
		}

		return v.Zond.HashTreeRoot()
	default:
		return zond.Root{}, ErrUnsupportedVersion
	}
}

// BodyRoot returns the body root of the blinded proposal.
func (v *VersionedBlindedProposal) BodyRoot() (zond.Root, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Body == nil {
			return zond.Root{}, ErrDataMissing
		}

		return v.Zond.Body.HashTreeRoot()
	default:
		return zond.Root{}, ErrUnsupportedVersion
	}
}

// ParentRoot returns the parent root of the blinded proposal.
func (v *VersionedBlindedProposal) ParentRoot() (zond.Root, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil {
			return zond.Root{}, ErrDataMissing
		}

		return v.Zond.ParentRoot, nil
	default:
		return zond.Root{}, ErrUnsupportedVersion
	}
}

// StateRoot returns the state root of the blinded proposal.
func (v *VersionedBlindedProposal) StateRoot() (zond.Root, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil {
			return zond.Root{}, ErrDataMissing
		}

		return v.Zond.StateRoot, nil
	default:
		return zond.Root{}, ErrUnsupportedVersion
	}
}

// TransactionsRoot returns the transactions root of the blinded proposal.
func (v *VersionedBlindedProposal) TransactionsRoot() (zond.Root, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Body == nil ||
			v.Zond.Body.ExecutionPayloadHeader == nil {
			return zond.Root{}, ErrDataMissing
		}

		return v.Zond.Body.ExecutionPayloadHeader.TransactionsRoot, nil
	default:
		return zond.Root{}, ErrUnsupportedVersion
	}
}

// FeeRecipient returns the fee recipient of the blinded proposal.
func (v *VersionedBlindedProposal) FeeRecipient() (zond.ExecutionAddress, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Body == nil ||
			v.Zond.Body.ExecutionPayloadHeader == nil {
			return zond.ExecutionAddress{}, ErrDataMissing
		}

		return v.Zond.Body.ExecutionPayloadHeader.FeeRecipient, nil
	default:
		return zond.ExecutionAddress{}, ErrUnsupportedVersion
	}
}

// Timestamp returns the timestamp of the blinded proposal.
func (v *VersionedBlindedProposal) Timestamp() (uint64, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Body == nil ||
			v.Zond.Body.ExecutionPayloadHeader == nil {
			return 0, ErrDataMissing
		}

		return v.Zond.Body.ExecutionPayloadHeader.Timestamp, nil
	default:
		return 0, ErrUnsupportedVersion
	}
}

// String returns a string version of the structure.
func (v *VersionedBlindedProposal) String() string {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil {
			return ""
		}

		return v.Zond.String()
	default:
		return "unknown version"
	}
}
