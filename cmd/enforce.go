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
	"fmt"

	"github.com/autonomy/conform/pkg/enforcer"
	"github.com/autonomy/conform/pkg/utilities"
	"github.com/spf13/cobra"
)

// enforceCmd represents the enforce command
var enforceCmd = &cobra.Command{
	Use:   "enforce",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			err := fmt.Errorf("The enforce command does not take arguments")

			return err
		}
		if err := utilities.CheckDockerVersion(); err != nil {
			return err
		}
		e, err := enforcer.New()
		if err != nil {
			return err
		}
		if err = e.Enforce(); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(enforceCmd)
}
