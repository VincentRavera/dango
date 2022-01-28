package cmd

import (
	"encoding/json"
	"fmt"
	"log"
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
		logger := log.Default()
		// cmd.Read
		rootpath, _ := utils.ReadEnviron(*logger)

		permissions := os.FileMode(0750)

		// Init git
		git.PlainInit(rootpath, false)
		os.Chdir(rootpath)
		// Gitignore
		gitignore ,e := os.Create(fmt.Sprintf("%s/.gitignore", rootpath))
		utils.ProcessSystemErrors(e, *logger)
		defer gitignore.Close()

		// Workspace
		gitignore.Write([]byte("workspace/"))
		os.MkdirAll(fmt.Sprintf("%s/workspace", rootpath), permissions)

		// current.json
		currentJsonPath := fmt.Sprintf("%s/current.json", rootpath)
		currentConfFile, e := os.Create(currentJsonPath)
		utils.ProcessSystemErrors(e, *logger)
		cf := data.CurrentConfig{
			Name: "main",
			Type: "development",
			Projects: []data.Project{},
		}
		cfbytes, e := json.Marshal(cf)
		utils.ProcessSystemErrors(e, *logger)
		currentConfFile.Write(cfbytes)

		return nil
	},
}
