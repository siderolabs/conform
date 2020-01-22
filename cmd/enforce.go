/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/talos-systems/conform/internal/enforcer"
	"github.com/talos-systems/conform/internal/policy"
)

// enforceCmd represents the enforce command
var enforceCmd = &cobra.Command{
	Use:   "enforce",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return errors.New("the enforce command does not take arguments")
		}
		// Done validating the arguments, do not print usage for errors
		// after this point
		cmd.SilenceUsage = true

		summarizer := cmd.Flags().Lookup("summary").Value.String()
		e, err := enforcer.New(summarizer)
		if err != nil {
			return fmt.Errorf("failed to create enforcer: %+v", err)
		}

		opts := []policy.Option{}

		if commitMsgFile := cmd.Flags().Lookup("commit-msg-file").Value.String(); commitMsgFile != "" {
			opts = append(opts, policy.WithCommitMsgFile(&commitMsgFile))
		}

		if err := e.Enforce(opts...); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	enforceCmd.Flags().String("commit-msg-file", "", "the path to the temporary commit message file")
	enforceCmd.Flags().String("summary", "none", "the summary method to use")
	rootCmd.AddCommand(enforceCmd)
}
