package conform

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/autonomy/conform/conform/config"
	"github.com/autonomy/conform/conform/git"
)

// Enforcer performs all the build actions for a rule.
type Enforcer struct {
	config  *config.Config
	rule    string
	GitInfo *git.Info
	Built   string
}

// NewEnforcer instantiates and returns an executer.
func NewEnforcer(rule string) (enforcer *Enforcer, err error) {
	enforcer = &Enforcer{}
	gitInfo, err := git.NewInfo()
	if err != nil {
		return
	}
	err = exportAll(gitInfo)
	if err != nil {
		return
	}
	date := []byte{}
	if gitInfo.IsTag {
		date, err = exec.Command("/bin/date").Output()
		if err != nil {
			return
		}
	}

	c, err := config.NewConfig()
	if err != nil {
		return
	}
	enforcer.config = c
	enforcer.GitInfo = gitInfo
	enforcer.Built = strings.TrimSuffix(string(date), "\n")
	enforcer.rule = rule

	return
}

// ExecuteBuild executes a docker build.
func (e *Enforcer) ExecuteBuild(dockerfile string) error {
	image := e.FormatImageNameSHA()
	if e.GitInfo.IsDirty {
		image = e.FormatImageNameDirty()
	}

	err := os.Setenv("CONFORM_IMAGE", image)
	if err != nil {
		return err
	}

	args := append([]string{"build", "--tag", image, "-f", "-", "."})
	command := exec.Command("docker", args...)
	command.Stdin = strings.NewReader(dockerfile)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	err = command.Start()
	if err != nil {
		return err
	}
	err = command.Wait()

	return err
}

// RenderDockerfile writes the final Dockerfile to disk.
func (e *Enforcer) RenderDockerfile(target *config.Rule) (dockerfile string, err error) {
	for _, p := range target.Templates {
		r, _err := e.RenderTemplate(p)
		if _err != nil {
			err = _err
			return
		}
		dockerfile = dockerfile + "\n" + *r
	}

	if e.config.Debug {
		fmt.Println(dockerfile)
	}

	return
}

// RenderTemplate executes the template and returns it.
func (e *Enforcer) RenderTemplate(s string) (*string, error) {
	if _s, ok := e.config.Templates[s]; ok {
		var wr bytes.Buffer
		tmpl, err := template.New("").Funcs(sprig.TxtFuncMap()).Parse(_s)
		if err != nil {
			return nil, err
		}

		err = tmpl.Execute(&wr, &e)
		if err != nil {
			return nil, err
		}

		str := wr.String()
		return &str, nil
	}

	return nil, fmt.Errorf("Template %q is not defined in conform.yaml", s)
}

// ExtractArtifact copies an artifact from a build.
func (e *Enforcer) ExtractArtifact(artifact string) error {
	return fmt.Errorf("Artifact %q is not defined in conform.yaml", artifact)
}

// ExecuteScript executes a script for a rule.
func (e *Enforcer) ExecuteScript(script string) error {
	if s, ok := e.config.Scripts[script]; ok {
		fmt.Printf("Running %q script\n", script)

		command := exec.Command("bash", "-c", s)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		err := command.Start()
		if err != nil {
			return err
		}
		err = command.Wait()
		if err != nil {
			return fmt.Errorf("Failed executing %q: %v", script, err)
		}

		return nil
	}

	return fmt.Errorf("Script %q is not defined in conform.yaml", script)
}

// ExecuteRule performs all the relevant actions specified in its' declaration.
func (e *Enforcer) ExecuteRule() error {
	if t, ok := e.config.Rules[e.rule]; ok {
		fmt.Printf("Enforcing %q\n", e.rule)
		for _, s := range t.Before {
			err := e.ExecuteScript(s)
			if err != nil {
				return err
			}
		}
		dockerfile, err := e.RenderDockerfile(t)
		if err != nil {
			return err
		}
		err = e.ExecuteBuild(dockerfile)
		if err != nil {
			return err
		}
		for _, s := range t.After {
			err := e.ExecuteScript(s)
			if err != nil {
				return err
			}
		}

		return nil
	}

	return fmt.Errorf("Rule %q is not defined in conform.yaml", e.rule)
}

func exportAll(gitInfo *git.Info) (err error) {
	fmt.Printf("Branch: %s\n", gitInfo.Branch)
	err = ExportConformVar("branch", gitInfo.Branch)
	if err != nil {
		return
	}
	fmt.Printf("SHA: %s\n", gitInfo.SHA)
	err = ExportConformVar("sha", gitInfo.SHA)
	if err != nil {
		return
	}
	fmt.Printf("Tag: %s\n", gitInfo.Tag)
	err = ExportConformVar("tag", gitInfo.Tag)
	if err != nil {
		return
	}
	fmt.Printf("IsTag: %s\n", strconv.FormatBool(gitInfo.IsTag))
	err = ExportConformVar("is_tag", strconv.FormatBool(gitInfo.IsTag))
	if err != nil {
		return
	}
	fmt.Printf("Prerelease: %s\n", gitInfo.Prerelease)
	err = ExportConformVar("prerelease", gitInfo.Prerelease)
	if err != nil {
		return
	}
	fmt.Printf("IsPrerelease: %s\n", strconv.FormatBool(gitInfo.IsPrerelease))
	err = ExportConformVar("is_prerelease", strconv.FormatBool(gitInfo.IsPrerelease))
	if err != nil {
		return
	}
	fmt.Printf("Status: \n%s\n", strings.TrimRight(gitInfo.Status, "\n"))
	err = ExportConformVar("status", strconv.FormatBool(gitInfo.IsDirty))
	if err != nil {
		return
	}
	fmt.Printf("IsDirty: %s\n", strconv.FormatBool(gitInfo.IsDirty))
	err = ExportConformVar("is_dirty", strconv.FormatBool(gitInfo.IsDirty))
	if err != nil {
		return
	}

	return
}

// ExportConformVar exports variable prefixed with CONFORM_
func ExportConformVar(name, value string) (err error) {
	variable := fmt.Sprintf("CONFORM_%s", strings.ToUpper(name))
	err = os.Setenv(variable, value)

	return
}

// FormatImageNameDirty formats the image name.
func (e *Enforcer) FormatImageNameDirty() string {
	return fmt.Sprintf("%s:%s", *e.config.Metadata.Repository, "dirty")
}

// FormatImageNameSHA formats the image name.
func (e *Enforcer) FormatImageNameSHA() string {
	return fmt.Sprintf("%s:%s", *e.config.Metadata.Repository, e.GitInfo.SHA)
}

// FormatImageNameTag formats the image name.
func (e *Enforcer) FormatImageNameTag() string {
	return fmt.Sprintf("%s:%s", *e.config.Metadata.Repository, e.GitInfo.Tag)
}

// FormatImageNameLatest formats the image name.
func (e *Enforcer) FormatImageNameLatest() string {
	return fmt.Sprintf("%s:%s", *e.config.Metadata.Repository, "latest")
}
