package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/VincentRavera/dango/data"
	"github.com/VincentRavera/dango/utils"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var Init = &cobra.Command{
	Use:    "init",
	Short:  "Initialize a dango project.",
	Long:   "Initialize an empty dango project.",
	Args:   cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// cmd.Read
		rootpath, _, _ := utils.ReadEnviron()

		permissions := os.FileMode(0750)

		// Init git
		git.PlainInit(rootpath, false)
		os.Chdir(rootpath)
		// Gitignore
		gitignore ,e := os.Create(fmt.Sprintf("%s/.gitignore", rootpath))
		utils.ProcessSystemErrors(e)
		defer gitignore.Close()

		// Workspace just in case
		gitignore.Write([]byte("workspace/"))
		os.MkdirAll(fmt.Sprintf("%s/workspace", rootpath), permissions)

		// current.json
		currentJsonPath := fmt.Sprintf("%s/current.json", rootpath)
		currentConfFile, e := os.Create(currentJsonPath)
		defer currentConfFile.Close()
		utils.ProcessSystemErrors(e)
		cf := data.CurrentConfig{
			Name: "main",
			Type: "development",
			Projects: []data.Project{},
		}
		cfbytes, e := json.Marshal(cf)
		utils.ProcessSystemErrors(e)
		currentConfFile.Write(cfbytes)

		return nil
	},
}
