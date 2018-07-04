package pipeline

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"

	"github.com/autonomy/conform/pkg/metadata"
	"github.com/autonomy/conform/pkg/release"
	"github.com/autonomy/conform/pkg/renderer"
	"github.com/autonomy/conform/pkg/service"
	"github.com/autonomy/conform/pkg/stage"
	"github.com/autonomy/conform/pkg/task"
	"github.com/mitchellh/mapstructure"
)

// Pipeline defines the stages and artifacts.
type Pipeline struct {
	Stages []string `yaml:"stages"`
}

// Build executes a docker build.
// nolint: gocyclo
func (p *Pipeline) Build(metadata *metadata.Metadata, stages map[string]*stage.Stage, tasks map[string]*task.Task) (err error) {
	if p == nil {
		return
	}
	for i, stageName := range p.Stages {
		if _, ok := stages[stageName]; !ok {
			return fmt.Errorf("Stage %q is not defined in conform.yaml", stageName)
		}
		stage := stages[stageName]
		// Anonymous func so the deferred service stop is executed at the end
		// of each stage.
		err = func() error {
			done := make(chan bool)
			for _, svc := range stage.Services {
				err = svc.Start()
				if err != nil {
					return err
				}
				defer svc.Stop()
				c := make(chan os.Signal, 1)
				signal.Notify(c, os.Interrupt)
				go func(s *service.Service) {
					for {
						select {
						case sig := <-c:
							fmt.Printf("Received %v signal, forcefully removing: %s\n", sig, s.Name)
							s.Rm()
						case <-done:
							return
						}
					}
				}(svc)
			}
			s, _err := p.render(metadata, stage.Tasks, tasks)
			if _err != nil {
				return _err
			}

			var image string
			if i+1 == len(p.Stages) {
				image = metadata.Docker.Image.Name + ":" + metadata.Docker.Image.Tag
			} else {
				image = metadata.Docker.Image.Name + ":" + stageName
			}

			_err = build(image, s)
			if _err != nil {
				return _err
			}
			for _, artifact := range stage.Artifacts {
				_err = p.extract(metadata.Git.SHA, image, artifact)
				if _err != nil {
					return _err
				}
			}

			return nil
		}()

		if err != nil {
			return err
		}
	}

	return nil
}

// Release executes the requested releases.
func (p *Pipeline) Release(metadata *metadata.Metadata, releases []*ReleaseDeclaration) (err error) {
	for _, r := range releases {
		if _, ok := releaseMap[r.Type]; !ok {
			return fmt.Errorf("Release %q is not defined", r.Type)
		}

		release := releaseMap[r.Type]
		err := mapstructure.Decode(r.Spec, release)
		if err != nil {
			return err
		}

		if err = release.Create(metadata); err != nil {
			return err
		}
		if err = release.Upload(); err != nil {
			return err
		}
	}

	return nil
}

// extract extracts an artifact from a docker image.
func (p *Pipeline) extract(sha, image string, artifact *stage.Artifact) error {
	argsSlice := [][]string{
		{"create", "--name=" + sha, image},
		{"cp", sha + ":" + artifact.Source, artifact.Destination},
		{"rm", sha},
	}
	for _, args := range argsSlice {
		command := exec.Command("docker", args...)
		command.Stderr = os.Stderr
		err := command.Start()
		if err != nil {
			return err
		}
		err = command.Wait()
		if err != nil {
			return err
		}
	}

	return nil
}

// render renders the stage tasks.
func (p *Pipeline) render(m *metadata.Metadata, requestedTasks []string, tasks map[string]*task.Task) (string, error) {
	var s string

	defer func(m *metadata.Metadata) {
		m.Docker.PreviousStage = ""
		m.Docker.NextStage = ""
		m.Docker.CurrentStage = ""
	}(m)

	for i, task := range requestedTasks {
		if _, ok := tasks[task]; !ok {
			return "", fmt.Errorf("Task %q is not defined in conform.yaml", task)
		}
		if i != 0 {
			m.Docker.PreviousStage = requestedTasks[i-1]
		}
		if i != len(requestedTasks)-1 {
			m.Docker.NextStage = requestedTasks[i+1]
		}
		m.Docker.CurrentStage = requestedTasks[i]
		rendered, err := renderer.RenderTemplate(tasks[task].Template, m)
		if err != nil {
			return "", err
		}
		s += rendered
	}

	return s, nil
}

func build(image, s string) error {
	args := append([]string{"build", "--network", "host", "--tag", image, "-f", "-", "."})
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

// ReleaseDeclaration allows a user to declare an arbitrary type along with a
// spec that will be decoded into the appropriate concrete type.
type ReleaseDeclaration struct {
	Type string      `yaml:"type"`
	Spec interface{} `yaml:"spec"`
}

// releaseMap defines the set of policies allowed within Conform.
var releaseMap = map[string]release.Release{
	"github": &release.GitHub{},
}
