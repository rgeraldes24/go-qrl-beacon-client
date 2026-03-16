// Copyright © 2021 - 2024 Attestant Limited.
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

import zond "github.com/theQRL/qrysm/v4/proto/prysm/v1alpha1"

// VersionedBeaconBlockBody contains a versioned beacon block body.
type VersionedBeaconBlockBody struct {
	Version DataVersion
	Zond    *zond.BeaconBlockBody
}

// String returns a string version of the structure.
func (v *VersionedBeaconBlockBody) String() string {
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
