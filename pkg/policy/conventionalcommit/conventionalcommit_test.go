package conventionalcommit

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"testing"

	"github.com/autonomy/conform/pkg/git"
	"github.com/autonomy/conform/pkg/metadata"
	"github.com/autonomy/conform/pkg/policy"
)

func RemoveAll(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}
}

func TestValidConventionalCommitPolicy(t *testing.T) {
	dir, err := ioutil.TempDir("", "test")
	if err != nil {
		log.Fatal(err)
	}
	defer RemoveAll(dir)
	err = os.Chdir(dir)
	if err != nil {
		t.Error(err)
	}
	err = initRepo()
	if err != nil {
		t.Error(err)
	}
	err = createValidCommit()
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

func TestInvalidConventionalCommitPolicy(t *testing.T) {
	dir, err := ioutil.TempDir("", "test")
	if err != nil {
		log.Fatal(err)
	}
	defer RemoveAll(dir)
	err = os.Chdir(dir)
	if err != nil {
		t.Error(err)
	}
	err = initRepo()
	if err != nil {
		t.Error(err)
	}
	err = createInvalidCommit()
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

func runCompliance() (*policy.Report, error) {
	g, err := git.NewGit()
	if err != nil {
		return nil, fmt.Errorf("failed to open git: %v", err)
	}
	message, err := g.Message()
	if err != nil {
		return nil, fmt.Errorf("failed to get commit message: %v", err)
	}
	c := &Conventional{}
	c.Types = []string{"type"}
	c.Scopes = []string{"scope"}
	m := &metadata.Metadata{
		Git: &metadata.Git{
			Message: message,
			IsClean: true,
		},
	}
	report := c.Compliance(m)

	return &report, nil
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
