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
	"os"
	"strings"

	"github.com/autonomy/conform/pkg/enforcer"
	"github.com/autonomy/conform/pkg/utilities"
	"github.com/spf13/cobra"
)

var (
	skipArray []string
	varArray  []string
	release   bool
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			err := fmt.Errorf("The build command does not take arguments")

			fmt.Println(err)
			os.Exit(1)
		}
		if err := utilities.CheckDockerVersion(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		e, err := enforcer.New(cmd.Flags())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, variable := range varArray {
			s := strings.Split(variable, "=")
			if len(s) != 2 {
				fmt.Printf("Variable key and value must be delimited by a '=': [%s]", variable)
				os.Exit(1)
			}
			e.Metadata.Variables[s[0]] = s[1]
		}
		for _, skip := range skipArray {
			for i, stage := range e.Pipeline.Stages {
				if stage == skip {
					e.Pipeline.Stages = append(e.Pipeline.Stages[:i], e.Pipeline.Stages[i+1:]...)
				}
			}
		}
		if err := e.Pipeline.Build(e.Metadata, e.Stages, e.Tasks); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if e.Script == nil {
			return
		}
		if err := e.Script.Execute(e.Metadata); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// if e.Metadata.Git.IsTag && release {
		if release {
			if err := e.Pipeline.Release(e.Metadata, e.Releases); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(buildCmd)
	buildCmd.Flags().StringArrayVar(&skipArray, "skip", []string{}, "skip a stage in the pipeline")
	buildCmd.Flags().StringArrayVar(&varArray, "var", []string{}, "set a variable")
	buildCmd.Flags().BoolVar(&release, "release", false, "perform a release")
}
