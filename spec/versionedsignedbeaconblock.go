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

	"github.com/rgeraldes24/go-qrl-beacon-client/spec/zond"
)

// VersionedSignedBeaconBlock contains a versioned signed beacon block.
type VersionedSignedBeaconBlock struct {
	Version DataVersion
	Zond    *zond.SignedBeaconBlock
}

// Slot returns the slot of the signed beacon block.
func (v *VersionedSignedBeaconBlock) Slot() (zond.Slot, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil || v.Zond.Message == nil {
			return 0, errors.New("no zond block")
		}

		return v.Zond.Message.Slot, nil
	default:
		return 0, errors.New("unknown version")
	}
}

// ProposerIndex returns the proposer index of the beacon block.
func (v *VersionedSignedBeaconBlock) ProposerIndex() (zond.ValidatorIndex, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil || v.Zond.Message == nil {
			return 0, errors.New("no zond block")
		}

		return v.Zond.Message.ProposerIndex, nil
	default:
		return 0, errors.New("unknown version")
	}
}

// ExecutionBlockHash returns the block hash of the beacon block.
func (v *VersionedSignedBeaconBlock) ExecutionBlockHash() (zond.Hash32, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Message == nil ||
			v.Zond.Message.Body == nil ||
			v.Zond.Message.Body.ExecutionPayload == nil {
			return zond.Hash32{}, errors.New("no zond block")
		}

		return v.Zond.Message.Body.ExecutionPayload.BlockHash, nil
	default:
		return zond.Hash32{}, errors.New("unknown version")
	}
}

// ExecutionBlockNumber returns the block number of the beacon block.
func (v *VersionedSignedBeaconBlock) ExecutionBlockNumber() (uint64, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Message == nil ||
			v.Zond.Message.Body == nil ||
			v.Zond.Message.Body.ExecutionPayload == nil {
			return 0, errors.New("no zond block")
		}

		return v.Zond.Message.Body.ExecutionPayload.BlockNumber, nil
	default:
		return 0, errors.New("unknown version")
	}
}

// ExecutionTransactions returns the execution payload transactions for the block.
func (v *VersionedSignedBeaconBlock) ExecutionTransactions() ([]zond.Transaction, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Message == nil ||
			v.Zond.Message.Body == nil ||
			v.Zond.Message.Body.ExecutionPayload == nil {
			return nil, errors.New("no zond block")
		}

		return v.Zond.Message.Body.ExecutionPayload.Transactions, nil
	default:
		return nil, errors.New("unknown version")
	}
}

// Graffiti returns the graffiti for the block.
func (v *VersionedSignedBeaconBlock) Graffiti() ([32]byte, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil || v.Zond.Message == nil || v.Zond.Message.Body == nil {
			return [32]byte{}, errors.New("no zond block")
		}

		return v.Zond.Message.Body.Graffiti, nil
	default:
		return [32]byte{}, errors.New("unknown version")
	}
}

// Attestations returns the attestations of the beacon block.
//
//nolint:gocyclo
func (v *VersionedSignedBeaconBlock) Attestations() ([]*VersionedAttestation, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil || v.Zond.Message == nil || v.Zond.Message.Body == nil {
			return nil, errors.New("no zond block")
		}

		versionedAttestations := make([]*VersionedAttestation, len(v.Zond.Message.Body.Attestations))
		for i, attestation := range v.Zond.Message.Body.Attestations {
			versionedAttestations[i] = &VersionedAttestation{
				Version: DataVersionZond,
				Zond:    attestation,
			}
		}

		return versionedAttestations, nil
	default:
		return nil, errors.New("unknown version")
	}
}

// Root returns the root of the beacon block.
func (v *VersionedSignedBeaconBlock) Root() (zond.Root, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil || v.Zond.Message == nil {
			return zond.Root{}, errors.New("no zond block")
		}

		return v.Zond.Message.HashTreeRoot()
	default:
		return zond.Root{}, errors.New("unknown version")
	}
}

// BodyRoot returns the body root of the beacon block.
func (v *VersionedSignedBeaconBlock) BodyRoot() (zond.Root, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil || v.Zond.Message == nil || v.Zond.Message.Body == nil {
			return zond.Root{}, errors.New("no zond block")
		}

		return v.Zond.Message.Body.HashTreeRoot()
	default:
		return zond.Root{}, errors.New("unknown version")
	}
}

// ParentRoot returns the parent root of the beacon block.
func (v *VersionedSignedBeaconBlock) ParentRoot() (zond.Root, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil || v.Zond.Message == nil {
			return zond.Root{}, errors.New("no zond block")
		}

		return v.Zond.Message.ParentRoot, nil
	default:
		return zond.Root{}, errors.New("unknown version")
	}
}

// StateRoot returns the state root of the beacon block.
func (v *VersionedSignedBeaconBlock) StateRoot() (zond.Root, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil || v.Zond.Message == nil {
			return zond.Root{}, errors.New("no zond block")
		}

		return v.Zond.Message.StateRoot, nil
	default:
		return zond.Root{}, errors.New("unknown version")
	}
}

// RandaoReveal returns the randao reveal of the beacon block.
func (v *VersionedSignedBeaconBlock) RandaoReveal() (zond.MLDSA87Signature, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil || v.Zond.Message == nil || v.Zond.Message.Body == nil {
			return zond.MLDSA87Signature{}, errors.New("no zond block")
		}

		return v.Zond.Message.Body.RANDAOReveal, nil
	default:
		return zond.MLDSA87Signature{}, errors.New("unknown version")
	}
}

// ExecutionData returns the execution data of the beacon block.
func (v *VersionedSignedBeaconBlock) ExecutionData() (*zond.ExecutionData, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil || v.Zond.Message == nil || v.Zond.Message.Body == nil {
			return nil, errors.New("no zond block")
		}

		return v.Zond.Message.Body.ExecutionData, nil
	default:
		return nil, errors.New("unknown version")
	}
}

// Deposits returns the deposits of the beacon block.
func (v *VersionedSignedBeaconBlock) Deposits() ([]*zond.Deposit, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil || v.Zond.Message == nil || v.Zond.Message.Body == nil {
			return nil, errors.New("no zond block")
		}

		return v.Zond.Message.Body.Deposits, nil
	default:
		return nil, errors.New("unknown version")
	}
}

// VoluntaryExits returns the voluntary exits of the beacon block.
func (v *VersionedSignedBeaconBlock) VoluntaryExits() ([]*zond.SignedVoluntaryExit, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil || v.Zond.Message == nil || v.Zond.Message.Body == nil {
			return nil, errors.New("no zond block")
		}

		return v.Zond.Message.Body.VoluntaryExits, nil
	default:
		return nil, errors.New("unknown version")
	}
}

// AttesterSlashings returns the attester slashings of the beacon block.
//
//nolint:gocyclo
func (v *VersionedSignedBeaconBlock) AttesterSlashings() ([]VersionedAttesterSlashing, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil || v.Zond.Message == nil || v.Zond.Message.Body == nil {
			return nil, errors.New("no zond block")
		}

		versionedAttesterSlashings := make([]VersionedAttesterSlashing, len(v.Zond.Message.Body.AttesterSlashings))
		for i, attesterSlashing := range v.Zond.Message.Body.AttesterSlashings {
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
func (v *VersionedSignedBeaconBlock) ProposerSlashings() ([]*zond.ProposerSlashing, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil || v.Zond.Message == nil || v.Zond.Message.Body == nil {
			return nil, errors.New("no zond block")
		}

		return v.Zond.Message.Body.ProposerSlashings, nil
	default:
		return nil, errors.New("unknown version")
	}
}

// SyncAggregate returns the sync aggregate of the beacon block.
func (v *VersionedSignedBeaconBlock) SyncAggregate() (*zond.SyncAggregate, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil || v.Zond.Message == nil || v.Zond.Message.Body == nil {
			return nil, errors.New("no zond block")
		}

		return v.Zond.Message.Body.SyncAggregate, nil
	default:
		return nil, errors.New("unknown version")
	}
}

// Withdrawals returns the withdrawals of the beacon block.
func (v *VersionedSignedBeaconBlock) Withdrawals() ([]*zond.Withdrawal, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil ||
			v.Zond.Message == nil ||
			v.Zond.Message.Body == nil ||
			v.Zond.Message.Body.ExecutionPayload == nil {
			return nil, errors.New("no zond block")
		}

		return v.Zond.Message.Body.ExecutionPayload.Withdrawals, nil
	default:
		return nil, errors.New("unknown version")
	}
}

// ExecutionPayload returns the execution payload of the signed beacon block.
func (v *VersionedSignedBeaconBlock) ExecutionPayload() (*VersionedExecutionPayload, error) {
	versionedExecutionPayload := &VersionedExecutionPayload{
		Version: v.Version,
	}

	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil || v.Zond.Message == nil || v.Zond.Message.Body == nil {
			return nil, errors.New("no zond block")
		}

		versionedExecutionPayload.Zond = v.Zond.Message.Body.ExecutionPayload
	default:
		return nil, errors.New("unknown version")
	}

	return versionedExecutionPayload, nil
}

// String returns a string version of the structure.
func (v *VersionedSignedBeaconBlock) String() string {
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
