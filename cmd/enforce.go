/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package cmd

import (
	"fmt"
	"os"

	"github.com/autonomy/conform/internal/enforcer"
	"github.com/autonomy/conform/internal/policy"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// enforceCmd represents the enforce command
var enforceCmd = &cobra.Command{
	Use:   "enforce",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			err := errors.Errorf("The enforce command does not take arguments")

			fmt.Println(err)
			os.Exit(1)
		}
		e, err := enforcer.New()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		opts := []policy.Option{}

		if commitMsgFile := cmd.Flags().Lookup("commit-msg-file").Value.String(); commitMsgFile != "" {
			opts = append(opts, policy.WithCommitMsgFile(&commitMsgFile))
		}

		e.Enforce(opts...)
	},
}

func init() {
	enforceCmd.Flags().String("commit-msg-file", "", "the path to the temporary commit message file")
	RootCmd.AddCommand(enforceCmd)
}
