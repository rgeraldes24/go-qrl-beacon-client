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

	zond "github.com/cyyber/qrysm/proto/qrysm/v1alpha1"
	"github.com/theQRL/go-qrl-beacon-client/api"
	"github.com/theQRL/go-qrl-beacon-client/spec"
)

// AttestationPool fetches the attestation pool for the given slot.
func (*Service) AttestationPool(_ context.Context,
	_ *api.AttestationPoolOpts,
) (
	*api.Response[[]*spec.VersionedAttestation],
	error,
) {
	data := make([]*spec.VersionedAttestation, 5)
	for i := range 5 {
		data[i] = &spec.VersionedAttestation{
			Version: spec.DataVersionZond,
			Zond: &zond.Attestation{
				Data: &zond.AttestationData{
					Source: &zond.Checkpoint{},
					Target: &zond.Checkpoint{},
				},
			},
		}
	}

	return &api.Response[[]*spec.VersionedAttestation]{
		Data:     data,
		Metadata: make(map[string]any),
	}, nil
}
