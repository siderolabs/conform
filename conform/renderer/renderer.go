package renderer

import (
	"bytes"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/autonomy/conform/conform/git"
	"github.com/autonomy/conform/conform/metadata"
)

// Renderer represents the rendering feature of Conform.
type Renderer struct{}

type data struct {
	Git   *git.Git
	Image string
	Built string
}

// Render parses a template.
func (r *Renderer) Render(m *metadata.Metadata, s string) (string, error) {
	var wr bytes.Buffer
	tmpl, err := template.New("").Funcs(sprig.TxtFuncMap()).Parse(s)
	if err != nil {
		return "", err
	}
	err = tmpl.Execute(&wr, m)
	if err != nil {
		return "", err
	}

	return wr.String(), nil
}
