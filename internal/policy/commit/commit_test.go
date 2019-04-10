/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package commit

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"testing"

	"github.com/autonomy/conform/internal/policy"
)

func RemoveAll(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}
}

func TestConventionalCommitPolicy(t *testing.T) {
	type testDesc struct {
		Name         string
		CreateCommit func() error
		ExpectValid  bool
	}

	for _, test := range []testDesc{
		{
			Name:         "Valid",
			CreateCommit: createValidCommit,
			ExpectValid:  true,
		},
		{
			Name:         "Invalid",
			CreateCommit: createInvalidCommit,
			ExpectValid:  false,
		},
		{
			Name:         "Empty",
			CreateCommit: createEmptyCommit,
			ExpectValid:  false,
		},
	} {
		func(test testDesc) {
			t.Run(test.Name, func(tt *testing.T) {
				dir, err := ioutil.TempDir("", "test")
				if err != nil {
					log.Fatal(err)
				}
				defer RemoveAll(dir)
				err = os.Chdir(dir)
				if err != nil {
					tt.Error(err)
				}
				err = initRepo()
				if err != nil {
					tt.Error(err)
				}

				err = test.CreateCommit()
				if err != nil {
					tt.Error(err)
				}
				report := runCompliance()

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

func runCompliance() *policy.Report {
	c := &Commit{
		Conventional: &Conventional{
			Types:  []string{"type"},
			Scopes: []string{"scope"},
		},
	}

	report := c.Compliance(&policy.Options{})

	return &report
}

func initRepo() error {
	_, err := exec.Command("git", "init").Output()
	if err != nil {
		return err
	}
	_, err = exec.Command("touch", "test").Output()
	if err != nil {
		return err
	}
	_, err = exec.Command("git", "add", "test").Output()

	return err
}

func createValidCommit() error {
	_, err := exec.Command("git", "-c", "user.name='test'", "-c", "user.email='test@autonomy.io'", "commit", "-m", "type(scope): description").Output()

	return err
}

func createInvalidCommit() error {
	_, err := exec.Command("git", "-c", "user.name='test'", "-c", "user.email='test@autonomy.io'", "commit", "-m", "invalid commit").Output()

	return err
}

func createEmptyCommit() error {
	_, err := exec.Command("git", "-c", "user.name='test'", "-c", "user.email='test@autonomy.io'", "commit", "--allow-empty-message", "-m", "").Output()

	return err
}
