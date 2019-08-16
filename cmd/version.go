/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package cmd

import (
	"bytes"
	"fmt"
	"html/template"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/talos-systems/conform/internal/constants"
)

var (
	shortVersion bool
	// Tag is set at build time.
	Tag string
	// SHA is set at build time.
	SHA string
	// Built is set at build time.
	Built string
)

const versionTemplate = constants.AppName + `:
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

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if shortVersion {
			PrintShortVersion()
		} else {
			PrintLongVersion()
		}
	},
}

func init() {
	versionCmd.Flags().BoolVar(&shortVersion, "short", false, "Print the short version")
	RootCmd.AddCommand(versionCmd)
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
	fmt.Println(fmt.Sprintf("%s %s-%s", constants.AppName, Tag, SHA))
}
