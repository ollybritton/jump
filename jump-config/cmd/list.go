package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list returns all directories that match a given alias.",
	Long:  `list will return all the paths to directories that a given alias match, separated by newlines.`,
	Run: func(cmd *cobra.Command, args []string) {
		listFiles(args[0])
	},
}

func listFiles(alias string) {
	matches, _ := FindMatches(alias)
	matchStrings := []string{}

	for key, val := range matches {
		matchStrings = append(matchStrings, key+"|"+val)
	}

	fmt.Println(strings.Join(matchStrings, ">"))

}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
