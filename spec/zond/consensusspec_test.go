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

package zond_test

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	ssz "github.com/ferranbt/fastssz"
	"github.com/goccy/go-yaml"
	"github.com/golang/snappy"
	clone "github.com/huandu/go-clone/generic"
	"github.com/rgeraldes24/go-qrl-beacon-client/spec/zond"
	require "github.com/stretchr/testify/require"
)

// TestConsensusSpec tests the types against the QRL consensus spec tests.
func TestConsensusSpec(t *testing.T) {
	if os.Getenv("CONSENSUS_SPEC_TESTS_DIR") == "" {
		t.Skip("CONSENSUS_SPEC_TESTS_DIR not supplied, not running spec tests")
	}

	tests := []struct {
		name string
		s    any
	}{
		{
			name: "AggregateAndProof",
			s:    &zond.AggregateAndProof{},
		},
		{
			name: "Attestation",
			s:    &zond.Attestation{},
		},
		{
			name: "AttestationData",
			s:    &zond.AttestationData{},
		},
		{
			name: "AttesterSlashing",
			s:    &zond.AttesterSlashing{},
		},
		{
			name: "BeaconBlock",
			s:    &zond.BeaconBlock{},
		},
		{
			name: "BeaconBlockBody",
			s:    &zond.BeaconBlockBody{},
		},
		{
			name: "BeaconBlockHeader",
			s:    &zond.BeaconBlockHeader{},
		},
		{
			name: "BeaconState",
			s:    &zond.BeaconState{},
		},
		{
			name: "Checkpoint",
			s:    &zond.Checkpoint{},
		},
		{
			name: "ContributionAndProof",
			s:    &zond.ContributionAndProof{},
		},
		{
			name: "Deposit",
			s:    &zond.Deposit{},
		},
		{
			name: "DepositData",
			s:    &zond.DepositData{},
		},
		{
			name: "DepositMessage",
			s:    &zond.DepositMessage{},
		},
		{
			name: "ExecutionData",
			s:    &zond.ExecutionData{},
		},
		{
			name: "ExecutionPayload",
			s:    &zond.ExecutionPayload{},
		},
		{
			name: "ExecutionPayloadHeader",
			s:    &zond.ExecutionPayloadHeader{},
		},
		{
			name: "Fork",
			s:    &zond.Fork{},
		},
		{
			name: "ForkData",
			s:    &zond.ForkData{},
		},
		{
			name: "HistoricalSummary",
			s:    &zond.HistoricalSummary{},
		},
		{
			name: "IndexedAttestation",
			s:    &zond.IndexedAttestation{},
		},
		{
			name: "PendingAttestation",
			s:    &zond.PendingAttestation{},
		},
		{
			name: "ProposerSlashing",
			s:    &zond.ProposerSlashing{},
		},
		{
			name: "SignedAggregateAndProof",
			s:    &zond.SignedAggregateAndProof{},
		},
		{
			name: "SignedBeaconBlock",
			s:    &zond.SignedBeaconBlock{},
		},
		{
			name: "SignedBeaconBlockHeader",
			s:    &zond.SignedBeaconBlockHeader{},
		},
		{
			name: "SignedContributionAndProof",
			s:    &zond.SignedContributionAndProof{},
		},
		{
			name: "SignedVoluntaryExit",
			s:    &zond.SignedVoluntaryExit{},
		},
		{
			name: "SyncAggregate",
			s:    &zond.SyncAggregate{},
		},
		{
			name: "SyncCommitteeContribution",
			s:    &zond.SyncCommitteeContribution{},
		},
		{
			name: "SyncCommitteeMessage",
			s:    &zond.SyncCommitteeMessage{},
		},
		{
			name: "Validator",
			s:    &zond.Validator{},
		},
		{
			name: "VoluntaryExit",
			s:    &zond.VoluntaryExit{},
		},
		{
			name: "Withdrawal",
			s:    &zond.Withdrawal{},
		},
	}

	baseDir := filepath.Join(os.Getenv("CONSENSUS_SPEC_TESTS_DIR"), "tests", "mainnet", "zond", "ssz_static")
	for _, test := range tests {
		dir := filepath.Join(baseDir, test.name, "ssz_random")
		require.NoError(t, filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if path == dir {
				// Only interested in subdirectories.
				return nil
			}
			require.NoError(t, err)
			if info.IsDir() {
				t.Run(fmt.Sprintf("%s/%s", test.name, info.Name()), func(t *testing.T) {
					s1 := clone.Clone(test.s)
					// Obtain the struct from the YAML.
					specYAML, err := os.ReadFile(filepath.Join(path, "value.yaml"))
					require.NoError(t, err)
					require.NoError(t, yaml.Unmarshal(specYAML, s1))
					// Confirm we can return to the YAML.
					remarshalledSpecYAML, err := yaml.Marshal(s1)
					require.NoError(t, err)
					require.YAMLEq(t, testYAMLFormat(specYAML), testYAMLFormat(remarshalledSpecYAML))

					// Obtain the struct from the SSZ.
					s2 := clone.Clone(test.s)
					compressedSpecSSZ, err := os.ReadFile(filepath.Join(path, "serialized.ssz_snappy"))
					require.NoError(t, err)
					var specSSZ []byte
					specSSZ, err = snappy.Decode(specSSZ, compressedSpecSSZ)
					require.NoError(t, err)
					require.NoError(t, s2.(ssz.Unmarshaler).UnmarshalSSZ(specSSZ))
					// Confirm we can return to the SSZ.
					remarshalledSpecSSZ, err := s2.(ssz.Marshaler).MarshalSSZ()
					require.NoError(t, err)
					require.Equal(t, specSSZ, remarshalledSpecSSZ)

					// Obtain the hash tree root from the YAML.
					specYAMLRoot, err := os.ReadFile(filepath.Join(path, "roots.yaml"))
					require.NoError(t, err)
					// Confirm we calculate the same root.
					generatedRootBytes, err := s2.(ssz.HashRoot).HashTreeRoot()
					require.NoError(t, err)
					generatedRoot := fmt.Sprintf("{root: '%#x'}\n", string(generatedRootBytes[:]))
					require.YAMLEq(t, string(specYAMLRoot), generatedRoot)
				})
			}

			return nil
		}))
	}
}

func testYAMLFormat(input []byte) string {
	val := make(map[string]any)
	if err := yaml.UnmarshalWithOptions(input, &val, yaml.UseOrderedMap()); err != nil {
		panic(err)
	}

	res, err := yaml.MarshalWithOptions(val, yaml.Flow(true))
	if err != nil {
		panic(err)
	}

	replacements := [][][]byte{
		{[]byte(`"`), []byte(`'`)},
		// Field 'extra_data' in BeaconBlockBody/case_3 has a non-standard format, fix here.
		{[]byte(`extra_data: 0,`), []byte(`extra_data: '0x',`)},
	}
	for _, replacement := range replacements {
		res = bytes.ReplaceAll(res, replacement[0], replacement[1])
	}

	return string(bytes.ToLower(res))
}
