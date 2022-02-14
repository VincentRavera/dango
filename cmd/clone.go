package cmd

import (
	"errors"
	"fmt"

	"github.com/VincentRavera/dango/utils"
	"github.com/spf13/cobra"
)

var Clone = &cobra.Command{
	Use:    "clone",
	Short:  "run git clone on projects",
	Long:   "Run clone on projects, update location",
	Args:   cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		rc := utils.GetConfig()
		if len(rc.Configuration.Projects) == 0 {
			return errors.New("No projects registered.\nUse: 'dango add' to populate your workspace.")
		}
		for i, project := range rc.Configuration.Projects {
			line, err := utils.CloneProject(i, project, rc)
			if err != nil {
				return err
			}
			if line != "" {
				fmt.Printf("%s\n", line)
			}
		}
		utils.SaveConfig()
		return nil
	},
}
