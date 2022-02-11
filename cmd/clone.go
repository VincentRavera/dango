package cmd

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/VincentRavera/dango/utils"
	"github.com/go-git/go-git/v5"
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
			// Test location if is already present
			isExists, err := utils.Exists(project.Location)
			if err == nil && isExists  {
				continue
			}
			// Clone
			fmt.Printf("Cloning %s\n", project.Name)
			newLocation := filepath.Join(rc.WorkPath, project.Name)
			_, err = git.PlainClone(newLocation, false, &git.CloneOptions{
				URL:      project.URL,
				RemoteName: project.Remote,
			})
			if err != nil {
				return fmt.Errorf("Cannot clone %s: %v", project.Name, err)
			}
			project.Location = newLocation
			// Update location
			utils.UpdateProject(i, project)

		}
		utils.SaveConfig()
		return nil
	},
}
