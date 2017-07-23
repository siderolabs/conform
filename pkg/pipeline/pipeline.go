package pipeline

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/autonomy/conform/pkg/metadata"
	"github.com/autonomy/conform/pkg/stage"
	"github.com/autonomy/conform/pkg/task"
)

// Pipeline defines the stages and artifacts.
type Pipeline struct {
	Stages []string `yaml:"stages"`
}

// Build executes a docker build.
func (p *Pipeline) Build(metadata *metadata.Metadata, stages map[string]*stage.Stage, tasks map[string]*task.Task) (err error) {
	for _, stage := range p.Stages {
		if _, ok := stages[stage]; !ok {
			return fmt.Errorf("Stage %q is not defined in conform.yaml", stage)
		}
		s, err := p.render(metadata, stages[stage].Tasks, tasks)
		if err != nil {
			return err
		}
		err = build(metadata.Docker.Image, s)
		if err != nil {
			return err
		}
		for _, artifact := range stages[stage].Artifacts {
			err = p.extract(metadata, artifact)
			if err != nil {
				return err
			}
		}
	}

	return
}

// extract extracts an artifact from a docker image.
func (p *Pipeline) extract(metadata *metadata.Metadata, artifact string) (err error) {
	argsSlice := [][]string{
		{"create", "--name=" + metadata.Git.SHA, metadata.Docker.Image},
		{"cp", metadata.Git.SHA + ":" + artifact, path.Base(artifact)},
		{"rm", metadata.Git.SHA},
	}
	for _, args := range argsSlice {
		command := exec.Command("docker", args...)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		err = command.Start()
		if err != nil {
			return err
		}
		err = command.Wait()
	}

	return
}

// render renders the stage tasks.
func (p *Pipeline) render(metadata *metadata.Metadata, requestedTasks []string, tasks map[string]*task.Task) (string, error) {
	var s string
	for _, task := range requestedTasks {
		if _, ok := tasks[task]; !ok {
			return "", fmt.Errorf("Task %q is not defined in conform.yaml", task)
		}
		err := tasks[task].Render(metadata)
		if err != nil {
			return "", err
		}
		s += tasks[task].Rendered
	}

	return s, nil
}

func build(image, s string) error {
	args := append([]string{"build", "--tag", image, "-f", "-", "."})
	command := exec.Command("docker", args...)
	command.Stdin = strings.NewReader(s)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err := command.Start()
	if err != nil {
		return err
	}

	return command.Wait()
}
