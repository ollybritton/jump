package cmd

import (
	"fmt"
	"regexp"

	"github.com/spf13/cast"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "returns the path associated with the alias",
	Long:  `parse will take in an alias for the directory and return the full path, as specified in the config file. If there are no matches, it returns NO_MATCH, and if there are multiple, it returns MULTIPLE_MATCH so that the bash function can decide what to do.`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("NO_ARGS")
			return
		}

		findDirectory(args[0])
	},
}

// FindMatches takes in a string and return
func FindMatches(s string) (map[string]string, []string) {
	aliases := viper.GetStringMap("aliases")
	regex := regexp.MustCompile("^" + s + "(.+)?")

	matches := map[string]string{}
	matchKeys := []string{}

	for key, val := range aliases {
		if regex.MatchString(key) {
			matches[string(key)] = cast.ToString(val)
			matchKeys = append(matchKeys, key)
		}
	}

	return matches, matchKeys

}

// findDirectory will either return the path associated with an alias or a code telling the associated bash function what to do.
func findDirectory(alias string) {
	matches, matchKeys := FindMatches(alias)

	switch len(matchKeys) {
	case 0:
		fmt.Println("NO_MATCH")
	case 1:
		fmt.Println(matches[matchKeys[0]])
	default:
		fmt.Println("MULTIPLE_MATCH")
	}
}

func init() {
	rootCmd.AddCommand(parseCmd)
}
