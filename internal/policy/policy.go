/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package policy

// Report summarizes the compliance of a policy.
type Report struct {
	Errors []error
}

// Policy is an interface that policies must implement.
type Policy interface {
	Compliance(*Options) Report
}

// Valid checks if a report is valid.
func (r Report) Valid() bool {
	return len(r.Errors) == 0
}
