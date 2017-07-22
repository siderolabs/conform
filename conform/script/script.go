package script

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/autonomy/conform/conform/metadata"
	"github.com/autonomy/conform/conform/renderer"
)

// Script defines a template that can be executed.
type Script struct {
	Template string `yaml:"template"`
	Rendered string
}

// Execute executes the pipeline script.
func (s *Script) Execute(metadata *metadata.Metadata) error {
	err := s.Render(metadata)
	if err != nil {
		return err
	}
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

// Render renders the script.
func (s *Script) Render(metadata *metadata.Metadata) error {
	renderer := renderer.Renderer{}
	rendered, err := renderer.Render(metadata, s.Template)
	if err != nil {
		return err
	}
	s.Rendered = rendered

	return nil
}
