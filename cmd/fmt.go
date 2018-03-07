// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
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
	"io/ioutil"
	"log"
	"sort"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

// fmtCmd represents the fmt command
var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		configBytes, err := ioutil.ReadFile(".conform.yaml")
		if err != nil {
			log.Fatal(err.Error())
		}
		m := yaml.MapSlice{}
		err = yaml.Unmarshal(configBytes, &m)
		if err != nil {
			log.Fatalf(err.Error())
		}
		for _, v := range m {
			if v.Key.(string) == "tasks" || v.Key.(string) == "stages" {
				mapSlice := v.Value.(yaml.MapSlice)
				sort.Slice(mapSlice, func(i, j int) bool {
					return mapSlice[i].Key.(string) < mapSlice[j].Key.(string)
				})
			}
		}
		configBytes, err = yaml.Marshal(m)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = ioutil.WriteFile(".conform.yaml", configBytes, 0644)
		if err != nil {
			log.Fatal(err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(fmtCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fmtCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fmtCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
