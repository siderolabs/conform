// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

//nolint:testpackage
package commit

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/siderolabs/conform/internal/policy"
)

//nolint:gocognit
func TestConventionalCommitPolicy(t *testing.T) {
	//nolint:govet
	type testDesc struct {
		Name         string
		CreateCommit func(t *testing.T) error
		ExpectValid  bool
		Conventional *Conventional
	}

	for _, test := range []testDesc{
		{
			Name:         "Valid",
			CreateCommit: createValidScopedCommit,
			ExpectValid:  true,
		},
		{
			Name:         "ValidBreaking",
			CreateCommit: createValidBreakingCommit,
			ExpectValid:  true,
		},
		{
			Name:         "InvalidBreakingSymbol",
			CreateCommit: createInvalidBreakingSymbolCommit,
			ExpectValid:  false,
		},
		{
			Name:         "ValidScopedBreaking",
			CreateCommit: createValidScopedBreakingCommit,
			ExpectValid:  true,
		},
		{
			Name:         "InvalidScopedBreaking",
			CreateCommit: createInvalidScopedBreakingCommit,
			ExpectValid:  false,
		},
		{
			Name:         "Invalid",
			CreateCommit: createInvalidCommit,
			ExpectValid:  false,
		},
		{
			Name:         "InvalidEmpty",
			CreateCommit: createInvalidEmptyCommit,
			ExpectValid:  false,
		},
		{
			Name:         "FixupRejected",
			CreateCommit: createFixupCommit,
			ExpectValid:  false,
		},
		{
			Name:         "FixupAccepted",
			CreateCommit: createFixupCommit,
			ExpectValid:  true,
			Conventional: acceptAutoSquashConventional(),
		},
		{
			Name:         "SquashAccepted",
			CreateCommit: createSquashCommit,
			ExpectValid:  true,
			Conventional: acceptAutoSquashConventional(),
		},
		{
			Name:         "AmendAccepted",
			CreateCommit: createAmendCommit,
			ExpectValid:  true,
			Conventional: acceptAutoSquashConventional(),
		},
	} {
		func(test testDesc) {
			t.Run(test.Name, func(tt *testing.T) {
				dir := t.TempDir()

				t.Chdir(dir)

				err := initRepo(t)
				if err != nil {
					tt.Error(err)
				}

				err = test.CreateCommit(tt)
				if err != nil {
					tt.Error(err)
				}

				report, err := runComplianceWithConventional(test.Conventional)
				if err != nil {
					t.Error(err)
				}

				if test.ExpectValid {
					if !report.Valid() {
						tt.Error("Report is invalid with valid conventional commit")
					}
				} else {
					if report.Valid() {
						tt.Error("Report is valid with invalid conventional commit")
					}
				}
			})
		}(test)
	}
}

func TestValidateDCO(t *testing.T) {
	type testDesc struct {
		Name          string
		CommitMessage string
		ExpectValid   bool
	}

	for _, test := range []testDesc{
		{
			Name:          "Valid DCO",
			CommitMessage: "something nice\n\nSigned-off-by: Foo Bar <foobar@example.org>\n\n",
			ExpectValid:   true,
		},
		{
			Name:          "Valid DCO with CRLF",
			CommitMessage: "something nice\r\n\r\nSigned-off-by: Foo Bar <foobar@example.org>\r\n\r\n",
			ExpectValid:   true,
		},
		{
			Name:          "No DCO",
			CommitMessage: "something nice\n\nnot signed\n",
			ExpectValid:   false,
		},
	} {
		// Fixes scopelint error.
		t.Run(test.Name, func(tt *testing.T) {
			var report policy.Report

			c := Commit{msg: test.CommitMessage}
			report.AddCheck(c.ValidateDCO())

			if test.ExpectValid {
				if !report.Valid() {
					tt.Error("Report is invalid with valid DCP")
				}
			} else {
				if report.Valid() {
					tt.Error("Report is valid with invalid DCO")
				}
			}
		})
	}
}

func TestValidConventionalCommitPolicy(t *testing.T) {
	dir := t.TempDir()

	t.Chdir(dir)

	err := initRepo(t)
	if err != nil {
		t.Error(err)
	}

	err = createValidScopedCommit(t)
	if err != nil {
		t.Error(err)
	}

	report, err := runCompliance()
	if err != nil {
		t.Error(err)
	}

	if !report.Valid() {
		t.Errorf("Report is invalid with valid conventional commit")
	}
}

func TestInvalidConventionalCommitPolicy(t *testing.T) {
	dir := t.TempDir()

	t.Chdir(dir)

	err := initRepo(t)
	if err != nil {
		t.Error(err)
	}

	err = createInvalidCommit(t)
	if err != nil {
		t.Error(err)
	}

	report, err := runCompliance()
	if err != nil {
		t.Error(err)
	}

	if report.Valid() {
		t.Errorf("Report is valid with invalid conventional commit")
	}
}

func TestEmptyConventionalCommitPolicy(t *testing.T) {
	dir := t.TempDir()

	t.Chdir(dir)

	err := initRepo(t)
	if err != nil {
		t.Error(err)
	}

	err = createInvalidEmptyCommit(t)
	if err != nil {
		t.Error(err)
	}

	report, err := runCompliance()
	if err != nil {
		t.Error(err)
	}

	if report.Valid() {
		t.Error("Report is valid with invalid conventional commit")
	}
}

func TestValidConventionalCommitPolicyRegex(t *testing.T) {
	dir := t.TempDir()

	t.Chdir(dir)

	err := initRepo(t)
	if err != nil {
		t.Error(err)
	}

	err = createValidCommitRegex(t)
	if err != nil {
		t.Error(err)
	}

	report, err := runCompliance()
	if err != nil {
		t.Error(err)
	}

	if !report.Valid() {
		t.Error("Report is invalid with valid conventional commit")
	}
}

func TestInvalidConventionalCommitPolicyRegex(t *testing.T) {
	dir := t.TempDir()

	t.Chdir(dir)

	err := initRepo(t)
	if err != nil {
		t.Error(err)
	}

	err = createInvalidCommitRegex(t)
	if err != nil {
		t.Error(err)
	}

	report, err := runCompliance()
	if err != nil {
		t.Error(err)
	}

	if report.Valid() {
		t.Error("Report is valid with invalid conventional commit")
	}
}

func TestValidRevisionRange(t *testing.T) {
	dir := t.TempDir()

	t.Chdir(dir)

	err := initRepo(t)
	if err != nil {
		t.Error(err)
	}

	revs, err := createValidCommitRange(t)
	if err != nil {
		t.Fatal(err)
	}

	// Test with a valid revision range
	report, err := runComplianceRange(revs[0], revs[len(revs)-1])
	if err != nil {
		t.Error(err)
	}

	if !report.Valid() {
		t.Error("Report is invalid with valid conventional commits")
	}

	// Test with HEAD as end of revision range
	report, err = runComplianceRange(revs[0], "HEAD")
	if err != nil {
		t.Error(err)
	}

	if !report.Valid() {
		t.Error("Report is invalid with valid conventional commits")
	}

	// Test with empty end of revision range (should fail)
	_, err = runComplianceRange(revs[0], "")
	if err == nil {
		t.Error("Invalid end of revision, got success, expecting failure")
	}

	// Test with empty start of revision (should fail)
	_, err = runComplianceRange("", "HEAD")
	if err == nil {
		t.Error("Invalid end of revision, got success, expecting failure")
	}

	// Test with start of revision not an ancestor of end of range (should fail)
	_, err = runComplianceRange(revs[1], revs[0])
	if err == nil {
		t.Error("Invalid end of revision, got success, expecting failure")
	}
}

func createValidCommitRange(t *testing.T) ([]string, error) {
	revs := make([]string, 0, 4)

	for i := range 4 {
		err := os.WriteFile("test", fmt.Append(nil, i), 0o644)
		if err != nil {
			return nil, fmt.Errorf("writing test file failed: %w", err)
		}

		_, err = exec.CommandContext(t.Context(), "git", "add", "test").Output()
		if err != nil {
			return nil, fmt.Errorf("git add failed: %w", err)
		}

		_, err = exec.CommandContext(t.Context(), "git", "-c", "user.name='test'", "-c", "user.email='test@siderolabs.io'", "commit", "-m", fmt.Sprintf("type(scope): description %d", i)).Output()
		if err != nil {
			return nil, fmt.Errorf("git commit failed: %w", err)
		}

		id, err := exec.CommandContext(t.Context(), "git", "rev-parse", "HEAD").Output()
		if err != nil {
			return nil, fmt.Errorf("rev-parse failed: %w", err)
		}

		revs = append(revs, strings.TrimSpace(string(id)))
	}

	return revs, nil
}

func runComplianceRange(id1, id2 string) (*policy.Report, error) {
	c := &Commit{
		Conventional: &Conventional{
			Types:  []string{"type"},
			Scopes: []string{"scope", "^valid"},
		},
	}

	return c.Compliance(&policy.Options{
		RevisionRange: fmt.Sprintf("%s..%s", id1, id2),
	})
}

func runCompliance() (*policy.Report, error) {
	return runComplianceWithConventional(nil)
}

func runComplianceWithConventional(conventional *Conventional) (*policy.Report, error) {
	if conventional == nil {
		conventional = &Conventional{
			Types:  []string{"type"},
			Scopes: []string{"scope", "^valid"},
		}
	}

	c := &Commit{Conventional: conventional}

	return c.Compliance(&policy.Options{})
}

func acceptAutoSquashConventional() *Conventional {
	return &Conventional{
		Types:            []string{"type"},
		Scopes:           []string{"scope", "^valid"},
		AcceptAutoSquash: true,
	}
}

func initRepo(t *testing.T) error {
	_, err := exec.CommandContext(t.Context(), "git", "init").Output()
	if err != nil {
		return err
	}

	_, err = exec.CommandContext(t.Context(), "touch", "test").Output()
	if err != nil {
		return err
	}

	_, err = exec.CommandContext(t.Context(), "git", "add", "test").Output()

	return err
}

func createValidScopedCommit(t *testing.T) error {
	_, err := exec.CommandContext(t.Context(), "git", "-c", "user.name='test'", "-c", "user.email='test@siderolabs.io'", "commit", "-m", "type(scope): description").Output()

	return err
}

func createValidBreakingCommit(t *testing.T) error {
	_, err := exec.CommandContext(t.Context(), "git", "-c", "user.name='test'", "-c", "user.email='test@siderolabs.io'", "commit", "-m", "feat!: description").Output()

	return err
}

func createInvalidBreakingSymbolCommit(t *testing.T) error {
	_, err := exec.CommandContext(t.Context(), "git", "-c", "user.name='test'", "-c", "user.email='test@siderolabs.io'", "commit", "-m", "feat$: description").Output()

	return err
}

func createValidScopedBreakingCommit(t *testing.T) error {
	_, err := exec.CommandContext(t.Context(), "git", "-c", "user.name='test'", "-c", "user.email='test@siderolabs.io'", "commit", "-m", "feat(scope)!: description").Output()

	return err
}

func createInvalidScopedBreakingCommit(t *testing.T) error {
	_, err := exec.CommandContext(t.Context(), "git", "-c", "user.name='test'", "-c", "user.email='test@siderolabs.io'", "commit", "-m", "feat!(scope): description").Output()

	return err
}

func createInvalidCommit(t *testing.T) error {
	_, err := exec.CommandContext(t.Context(), "git", "-c", "user.name='test'", "-c", "user.email='test@siderolabs.io'", "commit", "-m", "invalid commit").Output()

	return err
}

func createInvalidEmptyCommit(t *testing.T) error {
	_, err := exec.CommandContext(t.Context(), "git", "-c", "user.name='test'", "-c", "user.email='test@siderolabs.io'", "commit", "--allow-empty-message", "-m", "").Output()

	return err
}

func createValidCommitRegex(t *testing.T) error {
	_, err := exec.CommandContext(t.Context(), "git", "-c", "user.name='test'", "-c", "user.email='test@siderolabs.io'", "commit", "-m", "type(valid-1): description").Output()

	return err
}

func createInvalidCommitRegex(t *testing.T) error {
	_, err := exec.CommandContext(t.Context(), "git", "-c", "user.name='test'", "-c", "user.email='test@siderolabs.io'", "commit", "-m", "type(invalid-1): description").Output()

	return err
}

func createFixupCommit(t *testing.T) error {
	_, err := exec.CommandContext(t.Context(), "git", "-c", "user.name='test'", "-c", "user.email='test@siderolabs.io'", "commit", "-m", "fixup! deadbeef").Output()

	return err
}

func createSquashCommit(t *testing.T) error {
	_, err := exec.CommandContext(t.Context(), "git", "-c", "user.name='test'", "-c", "user.email='test@siderolabs.io'", "commit", "-m", "squash! deadbeef").Output()

	return err
}

func createAmendCommit(t *testing.T) error {
	_, err := exec.CommandContext(t.Context(), "git", "-c", "user.name='test'", "-c", "user.email='test@siderolabs.io'", "commit", "-m", "amend! deadbeef").Output()

	return err
}
