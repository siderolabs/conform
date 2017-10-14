package service

import (
	"log"
	"os"
	"os/exec"
)

// Service represents a container that can be run during a stage.
type Service struct {
	Name  string   `yaml:"name"`
	Image string   `yaml:"image"`
	Ports []string `yaml:"ports,omitempty"`
	Cmd   string   `yaml:"cmd,omitempty"`
}

// Start starts the service.
func (s *Service) Start() error {
	args := []string{"run", "--rm", "-d", "--network=host", "--name=" + s.Name, s.Image}
	if s.Cmd != "" {
		args = append(args, s.Cmd)
	}
	command := exec.Command("docker", args...)
	command.Stderr = os.Stderr
	err := command.Start()
	if err != nil {
		return err
	}
	err = command.Wait()

	return err
}

// Stop stops the service.
func (s *Service) Stop() {
	command := exec.Command("docker", []string{"stop", s.Name}...)
	command.Stderr = os.Stderr
	err := command.Start()
	if err != nil {
		log.Printf("Failed to stop service %q: %v", s.Name, err)
	}
	err = command.Wait()
	if err != nil {
		log.Printf("Failed to stop service %q: %v", s.Name, err)
	}
}

// Rm forcefully removes the service.
func (s *Service) Rm() {
	command := exec.Command("docker", []string{"rm", "-f", s.Name}...)
	command.Stderr = os.Stderr
	err := command.Start()
	if err != nil {
		log.Printf("Failed to remove service %q: %v", s.Name, err)
	}
	err = command.Wait()
	if err != nil {
		log.Printf("Failed to remove service %q: %v", s.Name, err)
	}
}
