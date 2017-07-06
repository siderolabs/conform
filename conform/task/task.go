package task

import (
	"github.com/autonomy/conform/conform/metadata"
	"github.com/autonomy/conform/conform/renderer"
)

// Task defines a stage that can be used within a pipeline.
type Task struct {
	Template string `yaml:"template"`
	Rendered string
}

// Render renders the stage.
func (s *Task) Render(metadata *metadata.Metadata) error {
	renderer := renderer.Renderer{}
	rendered, err := renderer.Render(metadata, s.Template)
	if err != nil {
		return err
	}
	s.Rendered = rendered

	return nil
}
