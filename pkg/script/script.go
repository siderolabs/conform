package script

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/autonomy/conform/pkg/metadata"
	"github.com/autonomy/conform/pkg/renderer"
)

// Script defines a template that can be executed.
type Script struct {
	Template string `yaml:"template"`
	Rendered string
}

// Execute executes the pipeline script.
func (s *Script) Execute(metadata *metadata.Metadata) error {
	r, err := renderer.RenderTemplate(s.Template, metadata)
	if err != nil {
		return err
	}
	s.Rendered = r
	command := exec.Command("bash", "-c", s.Rendered)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err = command.Start()
	if err != nil {
		return err
	}
	err = command.Wait()
	if err != nil {
		return fmt.Errorf("Failed executing script: %v", err)
	}

	return nil
}
