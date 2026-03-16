// Copyright © 2021 - 2025 Attestant Limited.
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
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"

	client "github.com/rgeraldes24/go-qrl-beacon-client"
	"github.com/rgeraldes24/go-qrl-beacon-client/api"
	"github.com/rgeraldes24/go-qrl-beacon-client/spec"
	"github.com/rgeraldes24/go-qrl-beacon-client/spec/zond"
)

// AttestationPool obtains the attestation pool for the given options.
func (s *Service) AttestationPool(ctx context.Context,
	opts *api.AttestationPoolOpts,
) (
	*api.Response[[]*spec.VersionedAttestation],
	error,
) {
	if err := s.assertIsSynced(ctx); err != nil {
		return nil, err
	}

	if opts == nil {
		return nil, client.ErrNoOptions
	}

	endpoint := "/qrl/v1/beacon/pool/attestations"

	queryItems := make([]string, 0)
	if opts.Slot != nil {
		queryItems = append(queryItems, fmt.Sprintf("slot=%d", *opts.Slot))
	}

	if opts.CommitteeIndex != nil {
		queryItems = append(queryItems, fmt.Sprintf("committee_index=%d", *opts.CommitteeIndex))
	}

	httpResponse, err := s.get(ctx, endpoint, strings.Join(queryItems, "&"), &opts.Common, false)
	if err != nil {
		return nil, err
	}

	switch httpResponse.contentType {
	case ContentTypeJSON:
		return s.attestationPoolFromJSON(ctx, opts, httpResponse)
	default:
		return nil, fmt.Errorf("unhandled content type %v", httpResponse.contentType)
	}
}

func (*Service) attestationPoolFromJSON(_ context.Context,
	opts *api.AttestationPoolOpts,
	httpResponse *httpResponse,
) (
	*api.Response[[]*spec.VersionedAttestation],
	error,
) {
	data, metadata, err := decodeJSONResponse(bytes.NewReader(httpResponse.body), []*spec.VersionedAttestation{})
	if err != nil {
		return nil, err
	}

	if err := verifyAttestationPool(opts, data); err != nil {
		return nil, err
	}

	return &api.Response[[]*spec.VersionedAttestation]{
		Metadata: metadata,
		Data:     data,
	}, nil
}

func verifyAttestationPool(opts *api.AttestationPoolOpts, data []*spec.VersionedAttestation) error {
	for _, datum := range data {
		switch datum.Version {
		case spec.DataVersionZond:
			if err := verifyPhase0Attestation(opts, datum.Zond); err != nil {
				return err
			}
		default:
			return errors.New("unsupported attestation version")
		}
	}

	return nil
}

func verifyPhase0Attestation(opts *api.AttestationPoolOpts, data *zond.Attestation) error {
	if opts.Slot != nil && data.Data.Slot != *opts.Slot {
		return errors.New("attestation data not for requested slot")
	}

	if opts.CommitteeIndex != nil && data.Data.Index != *opts.CommitteeIndex {
		return errors.New("attestation data not for requested committee index")
	}

	return nil
}
