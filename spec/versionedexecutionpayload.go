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

	"github.com/holiman/uint256"
	"github.com/theQRL/go-qrl-beacon-client/spec/zond"
)

// VersionedExecutionPayload contains a versioned execution payload.
type VersionedExecutionPayload struct {
	Version DataVersion
	Zond    *zond.ExecutionPayload
}

// IsEmpty returns true if there is no block.
func (v *VersionedExecutionPayload) IsEmpty() bool {
	return v.Zond == nil
}

// ParentHash returns the parent hash of the execution payload.
func (v *VersionedExecutionPayload) ParentHash() (zond.Hash32, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return zond.Hash32{}, errors.New("no zond execution payload")
		}

		return v.Zond.ParentHash, nil
	default:
		return zond.Hash32{}, errors.New("unknown version")
	}
}

// FeeRecipient returns the fee recipient of the execution payload.
func (v *VersionedExecutionPayload) FeeRecipient() (zond.ExecutionAddress, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return zond.ExecutionAddress{}, errors.New("no zond execution payload")
		}

		return v.Zond.FeeRecipient, nil
	default:
		return zond.ExecutionAddress{}, errors.New("unknown version")
	}
}

// StateRoot returns the state root of the execution payload.
func (v *VersionedExecutionPayload) StateRoot() (zond.Root, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return zond.Root{}, errors.New("no zond execution payload")
		}

		return v.Zond.StateRoot, nil
	default:
		return zond.Root{}, errors.New("unknown version")
	}
}

// ReceiptsRoot returns the receipts root of the execution payload.
func (v *VersionedExecutionPayload) ReceiptsRoot() (zond.Root, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return zond.Root{}, errors.New("no zond execution payload")
		}

		return v.Zond.ReceiptsRoot, nil
	default:
		return zond.Root{}, errors.New("unknown version")
	}
}

// LogsBloom returns the logs bloom of the execution payload.
func (v *VersionedExecutionPayload) LogsBloom() ([256]byte, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return [256]byte{}, errors.New("no zond execution payload")
		}

		return v.Zond.LogsBloom, nil
	default:
		return [256]byte{}, errors.New("unknown version")
	}
}

// PrevRandao returns the prev randao of the execution payload.
func (v *VersionedExecutionPayload) PrevRandao() ([32]byte, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return [32]byte{}, errors.New("no zond execution payload")
		}

		return v.Zond.PrevRandao, nil
	default:
		return [32]byte{}, errors.New("unknown version")
	}
}

// BlockNumber returns the block number of the execution payload.
func (v *VersionedExecutionPayload) BlockNumber() (uint64, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return 0, errors.New("no zond execution payload")
		}

		return v.Zond.BlockNumber, nil
	default:
		return 0, errors.New("unknown version")
	}
}

// GasLimit returns the gas limit of the execution payload.
func (v *VersionedExecutionPayload) GasLimit() (uint64, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return 0, errors.New("no zond execution payload")
		}

		return v.Zond.GasLimit, nil
	default:
		return 0, errors.New("unknown version")
	}
}

// GasUsed returns the gas used of the execution payload.
func (v *VersionedExecutionPayload) GasUsed() (uint64, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return 0, errors.New("no zond execution payload")
		}

		return v.Zond.GasUsed, nil
	default:
		return 0, errors.New("unknown version")
	}
}

// Timestamp returns the timestamp of the execution payload.
func (v *VersionedExecutionPayload) Timestamp() (uint64, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return 0, errors.New("no zond execution payload")
		}

		return v.Zond.Timestamp, nil
	default:
		return 0, errors.New("unknown version")
	}
}

// ExtraData returns the extra data of the execution payload.
func (v *VersionedExecutionPayload) ExtraData() ([]byte, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return nil, errors.New("no zond execution payload")
		}

		return v.Zond.ExtraData, nil
	default:
		return nil, errors.New("unknown version")
	}
}

// BaseFeePerGas returns the base fee per gas of the execution payload.
func (v *VersionedExecutionPayload) BaseFeePerGas() (*uint256.Int, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return nil, errors.New("no zond execution payload")
		}

		return uint256.NewInt(0).SetBytes(v.Zond.BaseFeePerGas[:]), nil
	default:
		return nil, errors.New("unknown version")
	}
}

// BlockHash returns the block hash of the execution payload.
func (v *VersionedExecutionPayload) BlockHash() (zond.Hash32, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return zond.Hash32{}, errors.New("no zond execution payload")
		}

		return v.Zond.BlockHash, nil
	default:
		return zond.Hash32{}, errors.New("unknown version")
	}
}

// Transactions returns the transactions of the execution payload.
func (v *VersionedExecutionPayload) Transactions() ([]zond.Transaction, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return nil, errors.New("no zond execution payload")
		}

		return v.Zond.Transactions, nil
	default:
		return nil, errors.New("unknown version")
	}
}

// Withdrawals returns the withdrawals of the execution payload.
func (v *VersionedExecutionPayload) Withdrawals() ([]*zond.Withdrawal, error) {
	switch v.Version {
	case DataVersionZond:
		if v.Zond == nil {
			return nil, errors.New("no zond execution payload")
		}

		return v.Zond.Withdrawals, nil
	default:
		return nil, errors.New("unknown version")
	}
}

// String returns a string version of the structure.
func (v *VersionedExecutionPayload) String() string {
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
