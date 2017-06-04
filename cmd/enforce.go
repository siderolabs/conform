// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/autonomy/conform/conform"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

// enforceCmd represents the enforce command
var enforceCmd = &cobra.Command{
	Use:   "enforce",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Invalid argument")
		}
		err := checkDockerVersion()
		if err != nil {
			return err
		}

		dp := conform.NewExecuter(args[0])
		err = dp.ExecuteRule()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(enforceCmd)
	RootCmd.Flags().BoolVar(&debug, "debug", false, "Debug rendering")
}

func checkDockerVersion() error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	serverVersion, err := cli.ServerVersion(context.Background())
	if err != nil {
		return err
	}
	minVersion, err := semver.NewVersion(minDockerVersion)
	if err != nil {
		return err
	}
	serverSemVer := semver.MustParse(serverVersion.Version)
	i := serverSemVer.Compare(minVersion)
	if i < 0 {
		return fmt.Errorf("At least Docker version %s is required", minDockerVersion)
	}

	return nil
}
