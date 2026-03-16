// Copyright © 2022 Attestant Limited.
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

package zond

import (
	"fmt"

	"github.com/goccy/go-yaml"
	"github.com/rgeraldes24/go-qrl-beacon-client/spec/zond"
)

// BlindedBeaconBlockBody represents the body of a blinded beacon block.
type BlindedBeaconBlockBody struct {
	RANDAOReveal           zond.MLDSA87Signature `ssz-size:"4627"`
	ExecutionData          *zond.ExecutionData
	Graffiti               [32]byte                    `ssz-size:"32"`
	ProposerSlashings      []*zond.ProposerSlashing    `ssz-max:"16"`
	AttesterSlashings      []*zond.AttesterSlashing    `ssz-max:"2"`
	Attestations           []*zond.Attestation         `ssz-max:"128"`
	Deposits               []*zond.Deposit             `ssz-max:"16"`
	VoluntaryExits         []*zond.SignedVoluntaryExit `ssz-max:"16"`
	SyncAggregate          *zond.SyncAggregate
	ExecutionPayloadHeader *zond.ExecutionPayloadHeader
}

// String returns a string version of the structure.
func (b *BlindedBeaconBlockBody) String() string {
	data, err := yaml.Marshal(b)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}
