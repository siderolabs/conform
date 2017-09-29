package renderer

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/autonomy/conform/pkg/metadata"
)

// Renderer renders all pipeline templates.
type Renderer struct{}

// Render renders a template.
func (r *Renderer) Render(m *metadata.Metadata, s string) (string, error) {
	var wr bytes.Buffer
	funcMap := makeFuncMap(m)
	tmpl, err := template.New("").Funcs(funcMap).Parse(s)
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

func makeFuncMap(m *metadata.Metadata) template.FuncMap {
	funcMap := sprig.TxtFuncMap()

	funcMap["fromURL"] = func(url string) (string, error) {
		resp, err := http.Get(url)
		if err != nil {
			return "", err
		}
		defer func() {
			_err := resp.Body.Close()
			if _err != nil {
				fmt.Printf("Failed to close the response body: %#v", _err)
			}
		}()
		if resp.StatusCode != http.StatusOK {
			return "", fmt.Errorf("Failed to download from URL: %d", resp.StatusCode)
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		rendered, err := RenderTemplate(string(b), m)
		if err != nil {
			return "", err
		}

		return rendered, nil
	}

	return funcMap
}
