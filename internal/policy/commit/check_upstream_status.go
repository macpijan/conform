// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package commit

import (
	"regexp"
	"strings"

	"github.com/pkg/errors"

	"github.com/3mdeb/conform/internal/policy"
)

// UpstreamStatusRegex is the regular expression used for Upstream Status.
var UpstreamStatusRegex = regexp.MustCompile(`^Upstream-Status:\s+([A-Za-z\s-]+)(?:\s*\[([^\]]+)\]|\s*\(([^)]+)\))?$`)

// UpstreamStatusCheck ensures that the commit message contains
// Upstream Status information.
type UpstreamStatusCheck struct {
	errors []error
}

// Name returns the name of the check.
func (d UpstreamStatusCheck) Name() string {
	return "UpstreamStatus"
}

// Message returns to check mes.sage.
func (d UpstreamStatusCheck) Message() string {
	if len(d.errors) != 0 {
		return d.errors[0].Error()
	}

	return "Upstream-Status was found"
}

// Errors returns any violations of the check.
func (d UpstreamStatusCheck) Errors() []error {
	return d.errors
}

// ValidateUpstreamStatus checks the commit message for Upstream Status.
func (c Commit) ValidateUpstreamStatus() policy.Check { //nolint:ireturn
	check := &UpstreamStatusCheck{}

	for _, line := range strings.Split(c.msg, "\n") {
		if UpstreamStatusRegex.MatchString(strings.TrimSpace(line)) {
			return check
		}
	}

	check.errors = append(check.errors, errors.Errorf("Commit does not have Upstream-Status information"))

	return check
}
