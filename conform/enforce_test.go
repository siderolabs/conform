package conform

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

var conformYAML = `
metadata:
  repository: test
scripts:
  before : |
    #!/bin/bash
    exit 0
  after : |
    #!/bin/bash
    exit 0
templates:
  test: |
    FROM scratch
rules:
  test:
    before:
      - before
    templates:
      - test
    after:
      - after
`

func TestNewEnforcer(t *testing.T) {
	dir, err := ioutil.TempDir("", "test")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)
	tmpfn := filepath.Join(dir, "conform.yaml")
	if err = ioutil.WriteFile(tmpfn, []byte(conformYAML), 0666); err != nil {
		log.Fatal(err)
	}
	os.Chdir(dir)
	output, err := exec.Command("git", "init").Output()
	if err != nil {
		t.Fatal(output)
	}
	output, err = exec.Command("git", "config", "--global", "user.email", "'test@autonomy.io'").Output()
	if err != nil {
		t.Fatal(output)
	}
	output, err = exec.Command("git", "config", "--global", "user.name", "'test'").Output()
	if err != nil {
		t.Fatal(output)
	}
	output, err = exec.Command("git", "add", "conform.yaml").Output()
	if err != nil {
		t.Fatal(output)
	}
	output, err = exec.Command("git", "commit", "-m", "'test'").Output()
	if err != nil {
		t.Fatal(output)
	}
	_, err = NewEnforcer("test")
	if err != nil {
		t.Fail()
	}
}
