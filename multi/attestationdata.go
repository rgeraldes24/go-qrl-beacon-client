// Copyright © 2021 Attestant Limited.
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

package multi

import (
	"context"

	zond "github.com/cyyber/qrysm/proto/qrysm/v1alpha1"
	consensusclient "github.com/theQRL/go-qrl-beacon-client"
	"github.com/theQRL/go-qrl-beacon-client/api"
)

// AttestationData fetches the attestation data for the given slot and committee index.
func (s *Service) AttestationData(ctx context.Context,
	opts *api.AttestationDataOpts,
) (
	*api.Response[*zond.AttestationData],
	error,
) {
	res, err := s.doCall(ctx, func(ctx context.Context, client consensusclient.Service) (any, error) {
		attestationData, err := client.(consensusclient.AttestationDataProvider).AttestationData(ctx, opts)
		if err != nil {
			return nil, err
		}

		return attestationData, nil
	}, nil)
	if err != nil {
		return nil, err
	}

	response, isResponse := res.(*api.Response[*zond.AttestationData])
	if !isResponse {
		return nil, ErrIncorrectType
	}

	return response, nil
}
