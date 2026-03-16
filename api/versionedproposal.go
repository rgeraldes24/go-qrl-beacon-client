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
	"math/big"

	"github.com/theQRL/go-qrl-beacon-client/spec"
)

// VersionedProposal contains a versioned proposal.
type VersionedProposal struct {
	Version        spec.DataVersion
	Blinded        bool
	ConsensusValue *big.Int
	ExecutionValue *big.Int
	Zond           *zond.BeaconBlock
	ZondBlinded    *apiv1zond.BlindedBeaconBlock
}

// IsEmpty returns true if there is no proposal.
func (v *VersionedProposal) IsEmpty() bool {
	return v.Zond == nil &&
		v.ZondBlinded == nil
}

// BodyRoot returns the body root of the proposal.
func (v *VersionedProposal) BodyRoot() (zond.Root, error) {
	if !v.bodyPresent() {
		return zond.Root{}, ErrDataMissing
	}

	switch v.Version {
	case spec.DataVersionZond:
		if v.Blinded {
			return v.ZondBlinded.Body.HashTreeRoot()
		}

		return v.Zond.Body.HashTreeRoot()
	default:
		return zond.Root{}, ErrUnsupportedVersion
	}
}

// ParentRoot returns the parent root of the proposal.
func (v *VersionedProposal) ParentRoot() (zond.Root, error) {
	if !v.proposalPresent() {
		return zond.Root{}, ErrDataMissing
	}

	switch v.Version {
	case spec.DataVersionZond:
		if v.Blinded {
			return v.ZondBlinded.ParentRoot, nil
		}

		return v.Zond.ParentRoot, nil
	default:
		return zond.Root{}, ErrUnsupportedVersion
	}
}

// ProposerIndex returns the proposer index of the proposal.
func (v *VersionedProposal) ProposerIndex() (zond.ValidatorIndex, error) {
	if !v.proposalPresent() {
		return 0, ErrDataMissing
	}

	switch v.Version {
	case spec.DataVersionZond:
		if v.Blinded {
			return v.ZondBlinded.ProposerIndex, nil
		}

		return v.Zond.ProposerIndex, nil
	default:
		return 0, ErrUnsupportedVersion
	}
}

// Root returns the root of the proposal.
func (v *VersionedProposal) Root() (zond.Root, error) {
	if !v.proposalPresent() {
		return zond.Root{}, ErrDataMissing
	}

	switch v.Version {
	case spec.DataVersionZond:
		if v.Blinded {
			return v.ZondBlinded.HashTreeRoot()
		}

		return v.Zond.HashTreeRoot()
	default:
		return zond.Root{}, ErrUnsupportedVersion
	}
}

// Slot returns the slot of the proposal.
func (v *VersionedProposal) Slot() (zond.Slot, error) {
	if !v.proposalPresent() {
		return 0, ErrDataMissing
	}

	switch v.Version {
	case spec.DataVersionZond:
		if v.Blinded {
			return v.ZondBlinded.Slot, nil
		}

		return v.Zond.Slot, nil
	default:
		return 0, ErrUnsupportedVersion
	}
}

// StateRoot returns the state root of the proposal.
func (v *VersionedProposal) StateRoot() (zond.Root, error) {
	if !v.proposalPresent() {
		return zond.Root{}, ErrDataMissing
	}

	switch v.Version {
	case spec.DataVersionZond:
		if v.Blinded {
			return v.ZondBlinded.StateRoot, nil
		}

		return v.Zond.StateRoot, nil
	default:
		return zond.Root{}, ErrUnsupportedVersion
	}
}

// Attestations returns the attestations of the proposal.
func (v *VersionedProposal) Attestations() ([]spec.VersionedAttestation, error) {
	if !v.bodyPresent() {
		return nil, ErrDataMissing
	}

	switch v.Version {
	case spec.DataVersionZond:
		if v.Blinded {
			versionedAttestations := make([]spec.VersionedAttestation, len(v.ZondBlinded.Body.Attestations))
			for i, attestation := range v.ZondBlinded.Body.Attestations {
				versionedAttestations[i] = spec.VersionedAttestation{
					Version: spec.DataVersionZond,
					Zond:    attestation,
				}
			}

			return versionedAttestations, nil
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

// Graffiti returns the graffiti of the proposal.
func (v *VersionedProposal) Graffiti() ([32]byte, error) {
	if !v.bodyPresent() {
		return [32]byte{}, ErrDataMissing
	}

	switch v.Version {
	case spec.DataVersionZond:
		if v.Blinded {
			return v.ZondBlinded.Body.Graffiti, nil
		}

		return v.Zond.Body.Graffiti, nil
	default:
		return [32]byte{}, ErrUnsupportedVersion
	}
}

// RandaoReveal returns the RANDAO reveal of the proposal.
func (v *VersionedProposal) RandaoReveal() (zond.MLDSA87Signature, error) {
	if !v.bodyPresent() {
		return zond.MLDSA87Signature{}, ErrDataMissing
	}

	switch v.Version {
	case spec.DataVersionZond:
		if v.Blinded {
			return v.ZondBlinded.Body.RANDAOReveal, nil
		}

		return v.Zond.Body.RANDAOReveal, nil
	default:
		return zond.MLDSA87Signature{}, ErrUnsupportedVersion
	}
}

// TODO(rgeraldes24)
/*
// Transactions returns the transactions of the proposal.
func (v *VersionedProposal) Transactions() ([]zond.Transaction, error) {
	if !v.payloadPresent() {
		return nil, ErrDataMissing
	}

	switch v.Version {
	case spec.DataVersionZond:
		if v.Blinded {
			return nil, ErrDataMissing
		}

		return v.Zond.Body.ExecutionPayload.Transactions, nil
	default:
		return nil, ErrUnsupportedVersion
	}
}
*/

// FeeRecipient returns the fee recipient of the proposal.
func (v *VersionedProposal) FeeRecipient() (zond.ExecutionAddress, error) {
	if !v.payloadPresent() {
		return zond.ExecutionAddress{}, ErrDataMissing
	}

	switch v.Version {
	case spec.DataVersionZond:
		if v.Blinded {
			return v.ZondBlinded.Body.ExecutionPayloadHeader.FeeRecipient, nil
		}

		return v.Zond.Body.ExecutionPayload.FeeRecipient, nil
	default:
		return zond.ExecutionAddress{}, ErrUnsupportedVersion
	}
}

// Timestamp returns the timestamp of the proposal.
func (v *VersionedProposal) Timestamp() (uint64, error) {
	if !v.payloadPresent() {
		return 0, ErrDataMissing
	}

	switch v.Version {
	case spec.DataVersionZond:
		if v.Blinded {
			return v.ZondBlinded.Body.ExecutionPayloadHeader.Timestamp, nil
		}

		return v.Zond.Body.ExecutionPayload.Timestamp, nil
	default:
		return 0, ErrUnsupportedVersion
	}
}

// GasLimit returns the gas limit of the proposal.
func (v *VersionedProposal) GasLimit() (uint64, error) {
	if !v.payloadPresent() {
		return 0, ErrDataMissing
	}

	switch v.Version {
	case spec.DataVersionZond:
		if v.Blinded {
			return v.ZondBlinded.Body.ExecutionPayloadHeader.GasLimit, nil
		}

		return v.Zond.Body.ExecutionPayload.GasLimit, nil
	default:
		return 0, ErrUnsupportedVersion
	}
}

// Value returns the value of the proposal, in Wei.
func (v *VersionedProposal) Value() *big.Int {
	value := big.NewInt(0)
	if v.ConsensusValue != nil {
		value = value.Add(value, v.ConsensusValue)
	}

	if v.ExecutionValue != nil {
		value = value.Add(value, v.ExecutionValue)
	}

	return value
}

// String returns a string version of the structure.
func (v *VersionedProposal) String() string {
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

func (v *VersionedProposal) proposalPresent() bool {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Blinded {
			return v.ZondBlinded != nil
		}

		return v.Zond != nil
	}

	return false
}

func (v *VersionedProposal) bodyPresent() bool {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Blinded {
			return v.ZondBlinded != nil && v.ZondBlinded.Body != nil
		}

		return v.Zond != nil && v.Zond.Body != nil
	}

	return false
}

//nolint:gocyclo // ignore
func (v *VersionedProposal) payloadPresent() bool {
	switch v.Version {
	case spec.DataVersionZond:
		if v.Blinded {
			return v.ZondBlinded != nil && v.ZondBlinded.Body != nil && v.ZondBlinded.Body.ExecutionPayloadHeader != nil
		}

		return v.Zond != nil && v.Zond.Body != nil && v.Zond.Body.ExecutionPayload != nil
	}

	return false
}
