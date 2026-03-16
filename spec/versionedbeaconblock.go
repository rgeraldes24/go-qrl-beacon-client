// Copyright © 2021 - 2024 Attestant Limited.
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

// VersionedBeaconBlock contains a versioned beacon block.
type VersionedBeaconBlock struct {
	Version DataVersion
	Zond    *zond.BeaconBlock
}

// IsEmpty returns true if there is no block.
func (v *VersionedBeaconBlock) IsEmpty() bool {
	return v.Zond == nil
}

// Slot returns the slot of the beacon block.
func (v *VersionedBeaconBlock) Slot() (zond.Slot, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return 0, errors.New("no zond block")
		}

		return v.Zond.Slot, nil
	default:
		return 0, errors.New("unknown version")
	}
}

// RandaoReveal returns the RANDAO reveal of the beacon block.
func (v *VersionedBeaconBlock) RandaoReveal() (zond.MLDSA87Signature, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return zond.MLDSA87Signature{}, errors.New("no zond block")
		}

		if v.Zond.Body == nil {
			return zond.MLDSA87Signature{}, errors.New("no zond block body")
		}

		return v.Zond.Body.RANDAOReveal, nil
	default:
		return zond.MLDSA87Signature{}, errors.New("unknown version")
	}
}

// Graffiti returns the graffiti of the beacon block.
func (v *VersionedBeaconBlock) Graffiti() ([32]byte, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return [32]byte{}, errors.New("no zond block")
		}

		if v.Zond.Body == nil {
			return [32]byte{}, errors.New("no zond block body")
		}

		return v.Zond.Body.Graffiti, nil
	default:
		return [32]byte{}, errors.New("unknown version")
	}
}

// ProposerIndex returns the proposer index of the beacon block.
func (v *VersionedBeaconBlock) ProposerIndex() (zond.ValidatorIndex, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return 0, errors.New("no zond block")
		}

		return v.Zond.ProposerIndex, nil
	default:
		return 0, errors.New("unknown version")
	}
}

// Root returns the root of the beacon block.
func (v *VersionedBeaconBlock) Root() (zond.Root, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return zond.Root{}, errors.New("no zond block")
		}

		return v.Zond.HashTreeRoot()
	default:
		return zond.Root{}, errors.New("unknown version")
	}
}

// BodyRoot returns the body root of the beacon block.
func (v *VersionedBeaconBlock) BodyRoot() (zond.Root, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return zond.Root{}, errors.New("no zond block")
		}

		if v.Zond.Body == nil {
			return zond.Root{}, errors.New("no zond block body")
		}

		return v.Zond.Body.HashTreeRoot()
	default:
		return zond.Root{}, errors.New("unknown version")
	}
}

// ParentRoot returns the parent root of the beacon block.
func (v *VersionedBeaconBlock) ParentRoot() (zond.Root, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return zond.Root{}, errors.New("no zond block")
		}

		return v.Zond.ParentRoot, nil
	default:
		return zond.Root{}, errors.New("unknown version")
	}
}

// StateRoot returns the state root of the beacon block.
func (v *VersionedBeaconBlock) StateRoot() (zond.Root, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return zond.Root{}, errors.New("no zond block")
		}

		return v.Zond.StateRoot, nil
	default:
		return zond.Root{}, errors.New("unknown version")
	}
}

// Attestations returns the attestations of the beacon block.
func (v *VersionedBeaconBlock) Attestations() ([]VersionedAttestation, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil || v.Zond.Body == nil {
			return nil, errors.New("no zond block")
		}

		versionedAttestations := make([]VersionedAttestation, len(v.Zond.Body.Attestations))
		for i, attestation := range v.Zond.Body.Attestations {
			versionedAttestations[i] = VersionedAttestation{
				Version: DataVersionZond,
				Zond:    attestation,
			}
		}

		return versionedAttestations, nil
	default:
		return nil, errors.New("unknown version")
	}
}

// AttesterSlashings returns the attester slashings of the beacon block.
func (v *VersionedBeaconBlock) AttesterSlashings() ([]VersionedAttesterSlashing, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil || v.Zond.Body == nil {
			return nil, errors.New("no zond block")
		}

		versionedAttesterSlashings := make([]VersionedAttesterSlashing, len(v.Zond.Body.AttesterSlashings))
		for i, attesterSlashing := range v.Zond.Body.AttesterSlashings {
			versionedAttesterSlashings[i] = VersionedAttesterSlashing{
				Version: DataVersionZond,
				Zond:    attesterSlashing,
			}
		}

		return versionedAttesterSlashings, nil
	default:
		return nil, errors.New("unknown version")
	}
}

// ProposerSlashings returns the proposer slashings of the beacon block.
func (v *VersionedBeaconBlock) ProposerSlashings() ([]*zond.ProposerSlashing, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil || v.Zond.Body == nil {
			return nil, errors.New("no zond block")
		}

		return v.Zond.Body.ProposerSlashings, nil
	default:
		return nil, errors.New("unknown version")
	}
}

// ExecutionPayload returns the execution payload of the beacon block.
func (v *VersionedBeaconBlock) ExecutionPayload() (*VersionedExecutionPayload, error) {
	versionedExecutionPayload := &VersionedExecutionPayload{
		Version: v.Version,
	}

	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil || v.Zond.Body == nil {
			return nil, errors.New("no zond block")
		}

		versionedExecutionPayload.Zond = v.Zond.Body.ExecutionPayload
	default:
		return nil, errors.New("unknown version")
	}

	return versionedExecutionPayload, nil
}

// String returns a string version of the structure.
func (v *VersionedBeaconBlock) String() string {
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
