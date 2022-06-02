// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/siderolabs/conform/internal/enforcer"
	"github.com/siderolabs/conform/internal/policy"
)

// enforceCmd represents the enforce command.
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

		reporter := cmd.Flags().Lookup("reporter").Value.String()
		e, err := enforcer.New(reporter)
		if err != nil {
			return fmt.Errorf("failed to create enforcer: %w", err)
		}

		opts := []policy.Option{}

		if commitMsgFile := cmd.Flags().Lookup("commit-msg-file").Value.String(); commitMsgFile != "" {
			opts = append(opts, policy.WithCommitMsgFile(&commitMsgFile))
		}

		if commitRef := cmd.Flags().Lookup("commit-ref").Value.String(); commitRef != "" {
			opts = append(opts, policy.WithCommitRef(commitRef))
		}

		return e.Enforce(opts...)
	},
}

func init() {
	enforceCmd.Flags().String("commit-msg-file", "", "the path to the temporary commit message file")
	enforceCmd.Flags().String("commit-ref", "", "the ref to compare git policies against")
	enforceCmd.Flags().String("reporter", "none", "the reporter method to use")
	rootCmd.AddCommand(enforceCmd)
}
