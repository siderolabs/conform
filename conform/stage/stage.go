package stage

import (
	"github.com/autonomy/conform/conform/metadata"
	"github.com/autonomy/conform/conform/renderer"
)

type Stage struct {
	Template string `yaml:"template"`
	Rendered string
}

func (s *Stage) Render(metadata *metadata.Metadata) error {
	renderer := renderer.Renderer{}
	rendered, err := renderer.Render(metadata, s.Template)
	if err != nil {
		return err
	}
	s.Rendered = rendered

	return nil
}
