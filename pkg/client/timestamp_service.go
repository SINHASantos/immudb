/*
Copyright 2025 Codenotary Inc. All rights reserved.

SPDX-License-Identifier: BUSL-1.1
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://mariadb.com/bsl11/

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"time"

	"github.com/codenotary/immudb/pkg/client/timestamp"
)

// TimestampService is a simple service returning the current time.
type TimestampService interface {
	GetTime() time.Time
}

type timestampService struct {
	ts timestamp.TsGenerator
}

// NewTimestampService creates new timestamp service returning current system time.
func NewTimestampService(ts timestamp.TsGenerator) TimestampService {
	return &timestampService{ts}
}

func (r *timestampService) GetTime() time.Time {
	return r.ts.Now()
}
