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
	"github.com/rgeraldes24/go-qrl-beacon-client/spec"
	"github.com/rgeraldes24/go-qrl-beacon-client/spec/zond"
)

// VersionedBlockRequest contains a versioned signed beacon block request.
type VersionedBlockRequest struct {
	Version spec.DataVersion
	Zond    *zond.SignedBeaconBlock
}

// Slot returns the slot of the signed beacon block.
func (v *VersionedBlockRequest) Slot() (zond.Slot, error) {
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

// ExecutionBlockHash returns the block hash of the beacon block.
func (v *VersionedBlockRequest) ExecutionBlockHash() (zond.Hash32, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Message == nil ||
			v.Zond.Message.Body == nil ||
			v.Zond.Message.Body.ExecutionPayload == nil {
			return zond.Hash32{}, ErrDataMissing
		}

		return v.Zond.Message.Body.ExecutionPayload.BlockHash, nil
	default:
		return zond.Hash32{}, ErrUnsupportedVersion
	}
}

// Attestations returns the attestations of the beacon block.
func (v *VersionedBlockRequest) Attestations() ([]spec.VersionedAttestation, error) {
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

// Root returns the root of the beacon block.
func (v *VersionedBlockRequest) Root() (zond.Root, error) {
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

// BodyRoot returns the body root of the beacon block.
func (v *VersionedBlockRequest) BodyRoot() (zond.Root, error) {
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

// ParentRoot returns the parent root of the beacon block.
func (v *VersionedBlockRequest) ParentRoot() (zond.Root, error) {
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

// StateRoot returns the state root of the beacon block.
func (v *VersionedBlockRequest) StateRoot() (zond.Root, error) {
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

// AttesterSlashings returns the attester slashings of the beacon block.
func (v *VersionedBlockRequest) AttesterSlashings() ([]spec.VersionedAttesterSlashing, error) {
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

// ProposerSlashings returns the proposer slashings of the beacon block.
func (v *VersionedBlockRequest) ProposerSlashings() ([]*zond.ProposerSlashing, error) {
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

// SyncAggregate returns the sync aggregate of the beacon block.
func (v *VersionedBlockRequest) SyncAggregate() (*zond.SyncAggregate, error) {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Message == nil ||
			v.Zond.Message.Body == nil {
			return nil, ErrDataMissing
		}

		return v.Zond.Message.Body.SyncAggregate, nil
	default:
		return nil, ErrUnsupportedVersion
	}
}

// String returns a string version of the structure.
func (v *VersionedBlockRequest) String() string {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil {
			return ""
		}

		return v.Zond.String()
	default:
		return "unsupported version"
	}
}
