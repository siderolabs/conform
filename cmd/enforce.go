/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/talos-systems/conform/internal/enforcer"
	"github.com/talos-systems/conform/internal/policy"
)

// enforceCmd represents the enforce command
var enforceCmd = &cobra.Command{
	Use:   "enforce",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("The enforce command does not take arguments")
			os.Exit(1)
		}
		summarizer := cmd.Flags().Lookup("summary").Value.String()
		e, err := enforcer.New(summarizer)
		if err != nil {
			log.Printf("failed to create enforcer: %+v\n", err)
			os.Exit(1)
		}

		opts := []policy.Option{}

		if commitMsgFile := cmd.Flags().Lookup("commit-msg-file").Value.String(); commitMsgFile != "" {
			opts = append(opts, policy.WithCommitMsgFile(&commitMsgFile))
		}

		if err := e.Enforce(opts...); err != nil {
			log.Printf("%+v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	enforceCmd.Flags().String("commit-msg-file", "", "the path to the temporary commit message file")
	enforceCmd.Flags().String("summary", "none", "the summary method to use")
	rootCmd.AddCommand(enforceCmd)
}
