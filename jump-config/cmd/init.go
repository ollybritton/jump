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
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Setup the program so that it can be used",
	Long:  `Add a bash/zsh function to config file such as '~/.zshrc' so that the jump command works as intended.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("You need to specify where to place the function code, for example")
			fmt.Println("jump-config init --language zsh $HOME/.zshrc")
			return
		}

		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			fmt.Println(args[0], "is not a valid path. Please try again with a valid path.")
			return
		}

		language, err := cmd.Flags().GetString("language")
		if err != nil {
			fmt.Println("Cannot find specified language. Use the --language flag to specify either bash or zsh.")
			return
		}

		if language != "bash" && language != "zsh" {
			fmt.Println("That is not a supported language. Try either 'zsh' or 'bash'.")
			return
		}

		addToConfig(args[0], language)
	},
}

func addToConfig(file string, language string) {
	existingConfig, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Error reading file, %v\n", err)
		return
	}

	if strings.Contains(string(existingConfig), "jump ()") {
		fmt.Println("It looks like a function named 'jump' already exists within that file.")
		return
	}

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	functionCodePath := path.Join(dir, language+".sh")
	functionCode, err := ioutil.ReadFile(functionCodePath)
	if err != nil {
		fmt.Println("Error reading function code in ", functionCodePath)
		return
	}

	newConfigString := string(existingConfig) + string("\n\n") + string(functionCode)
	newConfig := []byte(newConfigString)

	writeErr := ioutil.WriteFile(file, newConfig, 0644)
	if writeErr != nil {
		fmt.Println("Error writing function code to config.")
	}

}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringP("language", "l", "zsh", "What language to use for the shell function.")

}
