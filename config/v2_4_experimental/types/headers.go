// Copyright 2020 Red Hat, Inc.
//
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

package types

import (
	"github.com/coreos/ignition/config/shared/errors"
	"github.com/coreos/ignition/config/validate/report"
)

func (h HTTPHeaders) Validate() report.Report {
	for _, headerData := range h {
		// Validate that the header has just two elements
		if len(headerData) != 2 {
			return report.ReportFromError(errors.ErrInvalidHTTPHeader, report.EntryError)
		}

		// Header name can't be empty
		headerName := string(headerData[0])
		if headerName == "" {
			return report.ReportFromError(errors.ErrEmptyHTTPHeaderName, report.EntryError)
		}
	}

	// Validate that all header names in the list are unique
	set := make(map[string]struct{}) // New empty set
	for _, header := range h {
		set[string(header[0])] = struct{}{}
	}
	if len(set) != len(h) {
		return report.ReportFromError(errors.ErrDuplicateHTTPHeaders, report.EntryError)
	}

	return report.Report{}
}