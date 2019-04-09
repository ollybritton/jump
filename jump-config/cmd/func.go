// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"fmt"

	"github.com/spf13/cobra"
)

// funcCmd represents the func command
var funcCmd = &cobra.Command{
	Use:   "func",
	Short: "Output the function for the given options",
	Long:  `Print out the function associated with the given options, so you can add the function to your config file manually.`,
	Run: func(cmd *cobra.Command, args []string) {
		language, err := cmd.Flags().GetString("language")
		if err != nil {
			fmt.Println("Cannot find specified language. Use the --language flag to specify either bash or zsh.")
			return
		}

		if language != "bash" && language != "zsh" {
			fmt.Println("That is not a supported language. Try either 'zsh' or 'bash'.")
			return
		}

		color, err := cmd.Flags().GetBool("color")
		if err != nil {
			fmt.Println("There was a problem getting the color flag.")
			return
		}

		function, err := GetShellFunction(language, color)
		if err != nil {
			fmt.Printf("Error getting shell function: %v\n", err)
		}

		fmt.Println(function)

	},
}

func init() {
	rootCmd.AddCommand(funcCmd)

	funcCmd.Flags().StringP("language", "l", "zsh", "What language to use for the shell function.")
	funcCmd.Flags().BoolP("color", "c", false, "Use color for the command")
}
