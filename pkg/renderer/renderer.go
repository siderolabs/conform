package renderer

import (
	"bytes"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/autonomy/conform/pkg/metadata"
)

// Renderer renders all pipeline templates.
type Renderer struct{}

// Render renders a template.
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

// RenderTemplate renders a template with the provided metadata.
func RenderTemplate(contents string, metadata *metadata.Metadata) (string, error) {
	renderer := Renderer{}
	rendered, err := renderer.Render(metadata, contents)
	if err != nil {
		return "", err
	}

	return rendered, nil
}
