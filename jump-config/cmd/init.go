package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// initCmd represents the 'jump-config init' command, allowing for easily initialisation
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Setup the program so that it can be used",
	Long:  `Add a bash/zsh function to config file such as '~/.zshrc' so that the jump command works as intended.`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("You need to specify where to place the function code, for example")
			fmt.Println("jump-config init --language zsh ~/.zshrc")
			return
		}

		filepath := args[0]

		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			fmt.Println(filepath, "is not a valid path. Please try again with a valid path.")
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

		color, err := cmd.Flags().GetBool("color")
		if err != nil {
			fmt.Println("There was a problem getting the color flag.")
			return
		}

		AddToShellConfig(filepath, language, color)
	},
}

// CurrentDirectory will return the absolute path to the directory of the current file.
func CurrentDirectory() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Cannot find current directory")
	}

	dir := path.Dir(filename)
	return dir
}

// AppendToFile will append a string onto the end of a file
func AppendToFile(filename string, text string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		return err
	}

	return nil
}

// GetShellFunction will return the function that needs to be inserted into the user's shell config file in order to make the command work.
// "language" specifies the language, either "zsh" or "bash"
// "colors" specifies whether to use color output or not (boolean)
func GetShellFunction(language string, color bool) (string, error) {
	var filename string

	if color {
		filename = language + "_color.sh"
	} else {
		filename = language + ".sh"
	}

	filepath := path.Join(CurrentDirectory(), "../functions", filename)
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		return "", errors.New("Cannot read function file " + filepath)
	}

	return string(data), nil
}

// AddToShellConfig will add the required function to the shell config specified.
func AddToShellConfig(filepath string, language string, color bool) {
	function, err := GetShellFunction(language, color)
	if err != nil {
		panic(err)
	}

	existingConfig, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	if strings.Contains(string(existingConfig), "jump ()") {
		fmt.Println("It looks like a jump function already exists within " + filepath)
		return
	}

	err = AppendToFile(filepath, function)
	if err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringP("language", "l", "zsh", "What language to use for the shell function.")
	initCmd.Flags().BoolP("color", "c", false, "Use color for the command")

}
