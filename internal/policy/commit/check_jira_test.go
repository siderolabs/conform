/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package commit

import (
	"testing"
)

func TestCommit_ValidateJiraCheck(t *testing.T) {
	type fields struct {
		SpellCheck         *SpellCheck
		Conventional       *Conventional
		Header             *HeaderChecks
		Body               *BodyChecks
		DCO                bool
		GPG                bool
		MaximumOfOneCommit bool
		msg                string
	}

	type want struct {
		errorCount int
	}

	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{
			name: "Missing jira issue no type",
			fields: fields{
				Header: &HeaderChecks{
					Jira: &JiraChecks{
						Keys: []string{"JIRA", "PROJ"},
					},
				},
				msg: "invalid commit",
			},
			want: want{errorCount: 1},
		},
		{
			name: "Missing jira issue with type",
			fields: fields{
				Header: &HeaderChecks{
					Jira: &JiraChecks{
						Keys: []string{"JIRA", "PROJ"},
					},
				},
				msg: "fix: invalid commit",
			},
			want: want{errorCount: 1},
		},
		{
			name: "Valid commit",
			fields: fields{
				Header: &HeaderChecks{
					Jira: &JiraChecks{
						Keys: []string{"JIRA", "PROJ"},
					},
				},
				msg: "fix: [JIRA-1234] valid commit",
			},
			want: want{errorCount: 0},
		},
		{
			name: "Valid commit 2",
			fields: fields{
				Header: &HeaderChecks{
					Jira: &JiraChecks{
						Keys: []string{"JIRA", "PROJ"},
					},
				},
				msg: "fix: [PROJ-1234] valid commit",
			},
			want: want{errorCount: 0},
		},
		{
			name: "Invalid jira project",
			fields: fields{
				Header: &HeaderChecks{
					Jira: &JiraChecks{
						Keys: []string{"JIRA", "PROJ"},
					},
				},
				msg: "fix: [FALSE-1234] valid commit",
			},
			want: want{errorCount: 1},
		},
		{
			name: "Valid commit with scope",
			fields: fields{
				Header: &HeaderChecks{
					Jira: &JiraChecks{
						Keys: []string{"JIRA", "PROJ"},
					},
				},
				msg: "fix(test): [PROJ-1234] valid commit",
			},
			want: want{errorCount: 0},
		},
		{
			name: "Valid commit without square brackets",
			fields: fields{
				Header: &HeaderChecks{
					Jira: &JiraChecks{
						Keys: []string{"JIRA", "PROJ"},
					},
				},
				msg: "fix: PROJ-1234 valid commit",
			},
			want: want{errorCount: 0},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			c := Commit{
				SpellCheck:         tt.fields.SpellCheck,
				Conventional:       tt.fields.Conventional,
				Header:             tt.fields.Header,
				Body:               tt.fields.Body,
				DCO:                tt.fields.DCO,
				GPG:                tt.fields.GPG,
				MaximumOfOneCommit: tt.fields.MaximumOfOneCommit,
				msg:                tt.fields.msg,
			}
			got := c.ValidateJiraCheck()
			if len(got.Errors()) != tt.want.errorCount {
				t.Errorf("Wanted %d errors but got %d errors: %v", tt.want.errorCount, len(got.Errors()), got.Errors())
			}
		})
	}
}
