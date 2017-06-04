package conform

import (
	"bytes"
	"fmt"
	"runtime"
	"text/template"
)

var (
	// Tag is set at build time.
	Tag string
	// SHA is set at build time.
	SHA string
	// Built is set at build time.
	Built string
)

const versionTemplate = `Devise:
	Tag:         {{ .Tag }}
	SHA:         {{ .SHA }}
	Built:       {{ .Built }}
	Go version:  {{ .GoVersion }}
	OS/Arch:     {{ .Os }}/{{ .Arch }}
`

// Version contains verbose version information.
type Version struct {
	Tag       string
	SHA       string
	Built     string
	GoVersion string
	Os        string
	Arch      string
}

// PrintLongVersion prints verbose version information.
func PrintLongVersion() {
	v := Version{
		Tag:       Tag,
		SHA:       SHA,
		GoVersion: runtime.Version(),
		Os:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		Built:     Built,
	}

	var wr bytes.Buffer
	tmpl, err := template.New("version").Parse(versionTemplate)
	if err != nil {
		fmt.Println(err)
	}

	err = tmpl.Execute(&wr, v)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(wr.String())
}

// PrintShortVersion prints the tag and sha.
func PrintShortVersion() {
	fmt.Println(fmt.Sprintf("Devise %s-%s", Tag, SHA))
}
