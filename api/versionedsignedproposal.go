// Copyright © 2023, 2024 Attestant Limited.
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
	"errors"
	"math/big"

	apiv1zond "github.com/theQRL/go-qrl-beacon-client/api/v1/zond"
	"github.com/theQRL/go-qrl-beacon-client/spec"
	"github.com/theQRL/go-qrl-beacon-client/spec/zond"
)

// VersionedSignedProposal contains a versioned signed beacon node proposal.
type VersionedSignedProposal struct {
	Version        spec.DataVersion
	Blinded        bool
	ConsensusValue *big.Int
	ExecutionValue *big.Int
	Zond           *zond.SignedBeaconBlock
	ZondBlinded    *apiv1zond.SignedBlindedBeaconBlock
}

// AssertPresent throws an error if the expected proposal
// given the version and blinded fields is not present.
func (v *VersionedSignedProposal) AssertPresent() error {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Zond == nil && !v.Blinded {
			return errors.New("zond proposal not present")
		}

		if v.ZondBlinded == nil && v.Blinded {
			return errors.New("blinded zond proposal not present")
		}
	default:
		return errors.New("unsupported version")
	}

	return nil
}

// Slot returns the slot of the signed proposal.
func (v *VersionedSignedProposal) Slot() (zond.Slot, error) {
	err := v.assertMessagePresent()
	if err != nil {
		return 0, err
	}

	switch v.Version {
	case spec.DataVersionZond:
		if v.Blinded {
			return v.ZondBlinded.Message.Slot, nil
		}

		return v.Zond.Message.Slot, nil
	default:
		return 0, ErrUnsupportedVersion
	}
}

// ProposerIndex returns the proposer index of the signed proposal.
func (v *VersionedSignedProposal) ProposerIndex() (zond.ValidatorIndex, error) {
	if err := v.assertMessagePresent(); err != nil {
		return 0, err
	}

	switch v.Version {
	case spec.DataVersionZond:
		if v.Blinded {
			return v.ZondBlinded.Message.ProposerIndex, nil
		}

		return v.Zond.Message.ProposerIndex, nil
	default:
		return 0, ErrUnsupportedVersion
	}
}

// ExecutionBlockHash returns the hash of the execution payload.
func (v *VersionedSignedProposal) ExecutionBlockHash() (zond.Hash32, error) {
	if err := v.assertExecutionPayloadPresent(); err != nil {
		return zond.Hash32{}, err
	}

	switch v.Version {
	case spec.DataVersionZond:
		if v.Blinded {
			return v.ZondBlinded.Message.Body.ExecutionPayloadHeader.BlockHash, nil
		}

		return v.Zond.Message.Body.ExecutionPayload.BlockHash, nil
	default:
		return zond.Hash32{}, ErrUnsupportedVersion
	}
}

// String returns a string version of the structure.
func (v *VersionedSignedProposal) String() string {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Blinded {
			if v.ZondBlinded == nil {
				return ""
			}

			return v.ZondBlinded.String()
		}

		if v.Zond == nil {
			return ""
		}

		return v.Zond.String()
	default:
		return "unsupported version"
	}
}

// assertMessagePresent throws an error if the expected message
// given the version and blinded fields is not present.
//
//nolint:gocyclo // ignore
func (v *VersionedSignedProposal) assertMessagePresent() error {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Blinded {
			if v.ZondBlinded == nil ||
				v.ZondBlinded.Message == nil {
				return ErrDataMissing
			}
		} else {
			if v.Zond == nil ||
				v.Zond.Message == nil {
				return ErrDataMissing
			}
		}
	default:
		return ErrUnsupportedVersion
	}

	return nil
}

// assertExecutionPayloadPresent throws an error if the expected execution payload or payload header
// given the version and blinded fields is not present.
//
//nolint:gocyclo
func (v *VersionedSignedProposal) assertExecutionPayloadPresent() error {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Blinded {
			if v.ZondBlinded == nil ||
				v.ZondBlinded.Message == nil ||
				v.ZondBlinded.Message.Body == nil ||
				v.ZondBlinded.Message.Body.ExecutionPayloadHeader == nil {
				return ErrDataMissing
			}
		} else {
			if v.Zond == nil ||
				v.Zond.Message == nil ||
				v.Zond.Message.Body == nil ||
				v.Zond.Message.Body.ExecutionPayload == nil {
				return ErrDataMissing
			}
		}
	default:
		return ErrUnsupportedVersion
	}

	return nil
}
