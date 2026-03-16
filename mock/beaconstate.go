// Copyright © 2020 Attestant Limited.
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

package mock

import (
	"context"

	"github.com/theQRL/go-qrl-beacon-client/api"
	"github.com/theQRL/go-qrl-beacon-client/spec"
	"github.com/theQRL/go-qrl-beacon-client/spec/zond"
)

// BeaconState fetches a beacon state given a state ID.
func (s *Service) BeaconState(ctx context.Context,
	opts *api.BeaconStateOpts,
) (
	*api.Response[*spec.VersionedBeaconState],
	error,
) {
	if s.BeaconStateFunc != nil {
		return s.BeaconStateFunc(ctx, opts)
	}

	data := &spec.VersionedBeaconState{
		Version: spec.DataVersionZond,
		Zond: &zond.BeaconState{
			LatestBlockHeader:           &zond.BeaconBlockHeader{},
			ExecutionData:               &zond.ExecutionData{},
			PreviousJustifiedCheckpoint: &zond.Checkpoint{},
			CurrentJustifiedCheckpoint:  &zond.Checkpoint{},
			FinalizedCheckpoint:         &zond.Checkpoint{},
		},
	}

	return &api.Response[*spec.VersionedBeaconState]{
		Data:     data,
		Metadata: make(map[string]any),
	}, nil
}
