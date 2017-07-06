package pipeline

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/autonomy/conform/conform/metadata"
	"github.com/autonomy/conform/conform/stage"
)

type Pipeline struct {
	Stages []string `yaml:"stages"`
}

// Build executes a docker build.
func (p *Pipeline) Build(metadata *metadata.Metadata, stages map[string]*stage.Stage) (err error) {
	s, err := p.Render(metadata, stages)
	if err != nil {
		return err
	}
	args := append([]string{"build", "--tag", metadata.Docker.Image, "-f", "-", "."})
	command := exec.Command("docker", args...)
	command.Stdin = strings.NewReader(s)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err = command.Start()
	if err != nil {
		return err
	}
	err = command.Wait()

	return
}

func (p *Pipeline) Render(metadata *metadata.Metadata, stages map[string]*stage.Stage) (string, error) {
	var s string
	for _, stage := range p.Stages {
		if _, ok := stages[stage]; !ok {
			return "", fmt.Errorf("Stage %q not found", stage)
		}
		err := stages[stage].Render(metadata)
		if err != nil {
			return "", err
		}
		s += stages[stage].Rendered
	}

	return s, nil
}
