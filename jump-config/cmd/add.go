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
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an alias to the program",
	Long:  `Automatically add an alias to the program`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Two arguments are required, the name of the alias and the path to reference.")
			fmt.Println("Example: jump-config add [name] [path]")
			return
		}

		name := args[0]
		path := args[1]

		currentConfigPath := viper.ConfigFileUsed()
		currentConfig, err := ioutil.ReadFile(currentConfigPath)

		if err != nil {
			fmt.Println("Error reading current config file")
			return
		}

		var aliasCode string

		switch filepath.Ext(currentConfigPath) {
		case ".yaml":
			aliasCode = fmt.Sprintf("    %v: \"%v\"", name, path)
		case ".toml":
			aliasCode = fmt.Sprintf(`%v = "%v"`, name, path)
		}

		newConfig := []byte(string(currentConfig) + "\n" + aliasCode)
		err = ioutil.WriteFile(currentConfigPath, newConfig, os.ModePerm)

		if err != nil {
			fmt.Println("Error writing new config")
			return
		}

		fmt.Printf("Alias '%v' added to config.\n", name)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
