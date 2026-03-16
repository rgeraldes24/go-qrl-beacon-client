// Copyright © 2023 Attestant Limited.
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
	"github.com/theQRL/go-qrl-beacon-client/spec"
)

// VersionedSignedBlindedProposal contains a versioned signed blinded proposal.
type VersionedSignedBlindedProposal struct {
	Version spec.DataVersion
	Zond    *apiv1zond.SignedBlindedBeaconBlock
}

// Slot returns the slot of the signed blinded proposal.
func (v *VersionedSignedBlindedProposal) Slot() (zond.Slot, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Message == nil {
			return 0, ErrDataMissing
		}

		return v.Zond.Message.Slot, nil
	default:
		return 0, ErrUnsupportedVersion
	}
}

// Attestations returns the attestations of the signed blinded proposal.
func (v *VersionedSignedBlindedProposal) Attestations() ([]spec.VersionedAttestation, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Message == nil ||
			v.Zond.Message.Body == nil {
			return nil, ErrDataMissing
		}

		versionedAttestations := make([]spec.VersionedAttestation, len(v.Zond.Message.Body.Attestations))
		for i, attestation := range v.Zond.Message.Body.Attestations {
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
func (v *VersionedSignedBlindedProposal) Root() (zond.Root, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Message == nil {
			return zond.Root{}, ErrDataMissing
		}

		return v.Zond.Message.HashTreeRoot()
	default:
		return zond.Root{}, ErrUnsupportedVersion
	}
}

// BodyRoot returns the body root of the blinded proposal.
func (v *VersionedSignedBlindedProposal) BodyRoot() (zond.Root, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Message == nil ||
			v.Zond.Message.Body == nil {
			return zond.Root{}, ErrDataMissing
		}

		return v.Zond.Message.Body.HashTreeRoot()
	default:
		return zond.Root{}, ErrUnsupportedVersion
	}
}

// ParentRoot returns the parent root of the blinded proposal.
func (v *VersionedSignedBlindedProposal) ParentRoot() (zond.Root, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Message == nil {
			return zond.Root{}, ErrDataMissing
		}

		return v.Zond.Message.ParentRoot, nil
	default:
		return zond.Root{}, ErrUnsupportedVersion
	}
}

// StateRoot returns the state root of the blinded proposal.
func (v *VersionedSignedBlindedProposal) StateRoot() (zond.Root, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Message == nil {
			return zond.Root{}, ErrDataMissing
		}

		return v.Zond.Message.StateRoot, nil
	default:
		return zond.Root{}, ErrUnsupportedVersion
	}
}

// AttesterSlashings returns the attester slashings of the blinded proposal.
func (v *VersionedSignedBlindedProposal) AttesterSlashings() ([]spec.VersionedAttesterSlashing, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Message == nil ||
			v.Zond.Message.Body == nil {
			return nil, ErrDataMissing
		}

		versionedAttesterSlashings := make([]spec.VersionedAttesterSlashing, len(v.Zond.Message.Body.AttesterSlashings))
		for i, attesterSlashing := range v.Zond.Message.Body.AttesterSlashings {
			versionedAttesterSlashings[i] = spec.VersionedAttesterSlashing{
				Version: spec.DataVersionZond,
				Zond:    attesterSlashing,
			}
		}

		return versionedAttesterSlashings, nil
	default:
		return nil, ErrUnsupportedVersion
	}
}

// ProposerSlashings returns the proposer slashings of the blinded proposal.
func (v *VersionedSignedBlindedProposal) ProposerSlashings() ([]*zond.ProposerSlashing, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Message == nil ||
			v.Zond.Message.Body == nil {
			return nil, ErrDataMissing
		}

		return v.Zond.Message.Body.ProposerSlashings, nil
	default:
		return nil, ErrUnsupportedVersion
	}
}

// ProposerIndex returns the proposer index of the blinded proposal.
func (v *VersionedSignedBlindedProposal) ProposerIndex() (zond.ValidatorIndex, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Message == nil {
			return 0, ErrDataMissing
		}

		return v.Zond.Message.ProposerIndex, nil
	default:
		return 0, ErrUnsupportedVersion
	}
}

// ExecutionBlockHash returns the hash of the blinded proposal.
func (v *VersionedSignedBlindedProposal) ExecutionBlockHash() (zond.Hash32, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Message == nil ||
			v.Zond.Message.Body == nil ||
			v.Zond.Message.Body.ExecutionPayloadHeader == nil {
			return zond.Hash32{}, ErrDataMissing
		}

		return v.Zond.Message.Body.ExecutionPayloadHeader.BlockHash, nil
	default:
		return zond.Hash32{}, ErrUnsupportedVersion
	}
}

// ExecutionBlockNumber returns the block number of the blinded proposal.
func (v *VersionedSignedBlindedProposal) ExecutionBlockNumber() (uint64, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Message == nil ||
			v.Zond.Message.Body == nil ||
			v.Zond.Message.Body.ExecutionPayloadHeader == nil {
			return 0, ErrDataMissing
		}

		return v.Zond.Message.Body.ExecutionPayloadHeader.BlockNumber, nil
	default:
		return 0, ErrUnsupportedVersion
	}
}

// Signature returns the signature of the blinded proposal.
func (v *VersionedSignedBlindedProposal) Signature() (zond.MLDSA87Signature, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil {
			return zond.MLDSA87Signature{}, ErrDataMissing
		}

		return v.Zond.Signature, nil
	default:
		return zond.MLDSA87Signature{}, ErrUnsupportedVersion
	}
}
