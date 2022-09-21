// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"errors"
	"fmt"

	git "github.com/go-git/go-git/v5"
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
		} else {
			mainBranch, err := detectMainBranch()
			if err != nil {
				return fmt.Errorf("failed to detect main branch: %w", err)
			}
			if mainBranch != "" {
				opts = append(opts, policy.WithCommitRef(fmt.Sprintf("refs/heads/%s", mainBranch)))
			}
		}

		if baseBranch := cmd.Flags().Lookup("base-branch").Value.String(); baseBranch != "" {
			opts = append(opts, policy.WithRevisionRange(fmt.Sprintf("%s..HEAD", baseBranch)))
		} else if revisionRange := cmd.Flags().Lookup("revision-range").Value.String(); revisionRange != "" {
			opts = append(opts, policy.WithRevisionRange(revisionRange))
		}

		return e.Enforce(opts...)
	},
}

func init() {
	enforceCmd.Flags().String("commit-msg-file", "", "the path to the temporary commit message file")
	enforceCmd.Flags().String("commit-ref", "", "the ref to compare git policies against")
	enforceCmd.Flags().String("reporter", "none", "the reporter method to use")
	enforceCmd.Flags().String("revision-range", "", "<commit1>..<commit2>")
	enforceCmd.Flags().String("base-branch", "", "base branch to compare with")
	rootCmd.AddCommand(enforceCmd)
}

func detectMainBranch() (string, error) {
	mainBranch := "main"

	repo, err := git.PlainOpen(".")
	if err != nil {
		// not a git repo, ignore
		return "", nil //nolint:nilerr
	}

	c, err := repo.Config()
	if err != nil {
		return "", fmt.Errorf("failed to get repository configuration: %w", err)
	}

	rawConfig := c.Raw

	const branchSectionName = "branch"

	branchSection := rawConfig.Section(branchSectionName)
	for _, b := range branchSection.Subsections {
		remote := b.Option("remote")
		if remote == git.DefaultRemoteName {
			mainBranch = b.Name

			break
		}
	}

	return mainBranch, nil
}
