// Copyright © 2020, 2021 Attestant Limited.
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

package http

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/theQRL/go-qrl-beacon-client/api"
	"github.com/theQRL/go-qrl-beacon-client/spec/zond"
)

// SlotFromStateID parses the state ID and returns the relevant slot.
func (*Service) SlotFromStateID(_ context.Context, stateID string) (zond.Slot, error) {
	var slot zond.Slot

	switch {
	case stateID == "genesis":
		slot = 0
	case stateID == "justified":
		return 0, errors.New("state from justified not implemented")
	case stateID == "finalized":
		return 0, errors.New("state from finalized not implemented")
	case stateID == "head":
		return 0, errors.New("state from head not implemented")
	case strings.HasPrefix(stateID, "0x"):
		return 0, errors.New("state from state root not implemented")
	default:
		// State ID should be a slot.
		tmp, err := strconv.ParseUint(stateID, 10, 64)
		if err != nil {
			return 0, errors.Join(fmt.Errorf("failed to parse state %s as a slot", stateID), err)
		}

		slot = zond.Slot(tmp)
	}

	return slot, nil
}

// EpochFromStateID parses the state ID and returns the relevant epoch.
func (s *Service) EpochFromStateID(ctx context.Context, stateID string) (zond.Epoch, error) {
	var epoch zond.Epoch

	switch {
	case stateID == "genesis":
		epoch = 0
	case stateID == "justified":
		response, err := s.Finality(ctx, &api.FinalityOpts{State: stateID})
		if err != nil {
			return 0, errors.Join(errors.New("failed to obtain finality for justified epoch"), err)
		}

		epoch = response.Data.Justified.Epoch
	case stateID == "finalized":
		response, err := s.Finality(ctx, &api.FinalityOpts{State: stateID})
		if err != nil {
			return 0, errors.Join(errors.New("failed to obtain finality for finalized epoch"), err)
		}

		epoch = response.Data.Justified.Epoch
	case stateID == "head":
		return 0, errors.New("epoch from head not implemented")
	case strings.HasPrefix(stateID, "0x"):
		return 0, errors.New("epoch from state root not implemented")
	default:
		// State ID should be a slot.
		tmp, err := strconv.ParseUint(stateID, 10, 64)
		if err != nil {
			return 0, errors.Join(fmt.Errorf("failed to parse state %s as a slot", stateID), err)
		}

		slotsPerEpoch, err := s.SlotsPerEpoch(ctx)
		if err != nil {
			return 0, errors.Join(errors.New("failed to obtain slots per epoch"), err)
		}

		epoch = zond.Epoch(tmp / slotsPerEpoch)
	}

	return epoch, nil
}
