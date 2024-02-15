// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/siderolabs/conform/internal/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints Kres version.",
	Long:  `Prints Kres version.`,
	Args:  cobra.NoArgs,
	Run: func(*cobra.Command, []string) {
		line := fmt.Sprintf("%s version %s (%s)", version.Name, version.Tag, version.SHA)
		fmt.Println(line)
	},
}
