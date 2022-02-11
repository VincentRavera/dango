package cmd

import (
	"errors"
	"fmt"

	"github.com/VincentRavera/dango/utils"
	"github.com/spf13/cobra"
)

var Config = &cobra.Command{
	Use:    "config",
	Short:  "read config",
	Long:   "Ensure config is readable.",
	Args:   cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		rc := utils.GetConfig()
		fmt.Println(rc.String())
		if len(rc.Configuration.Projects) == 0 {
			return errors.New("No projects registered.\nUse: 'dango add' to populate your workspace.")
		}
		return nil
	},
}
