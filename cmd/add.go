package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/VincentRavera/dango/utils"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var Add = &cobra.Command{
	Use: 	"add",
	Short:	"Add a project to manage.",
	Long: 	"Add a project or git url to project.",
	Args: 	func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("A path or git URL is required.")
		}
		for _, value := range args {
			if ! validateArg(value) {
				return errors.New(fmt.Sprintf("The argument %s is not valid", value))
			}
		}
		return nil
	},
	Run: 	func(cmd *cobra.Command, args []string) {
		l := log.Default()
		rootConf := utils.GetConfig(*l)
		for _, value := range args {
			// works for path and urls
			projectName := filepath.Base(value)
			fmt.Println(projectName)

			path := value
			if strings.HasPrefix(value, "http") || strings.HasPrefix(value, "git@") {
				path = rootConf.WorkPath + projectName
				git.PlainClone(path, false, nil)
			}

			utils.AddProject(utils.ScanPath(path), rootConf)
		}
	},
}

// https://stackoverflow.com/a/10510783
// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return false, err
}
func validateArg(arg string) bool {
	if len(arg) < 0 {
		return false
	}
	if strings.HasPrefix(arg, "/") {
		isargExisting, _ := exists(arg)
		if isargExisting {
			return true
		} else {
			return false
		}
	}
	if strings.HasPrefix(arg, "http") || strings.HasPrefix(arg, "git@") {
		return true
	}
	return false

}
