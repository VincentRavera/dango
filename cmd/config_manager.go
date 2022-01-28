package cmd

import (
	"log"

	"github.com/VincentRavera/dango/utils"
	"github.com/spf13/cobra"
)

var Config = &cobra.Command{
	Use:    "config",
	Short:  "read config",
	Long:   "Ensure config is readable.",
	Args:   cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := log.Default()
		utils.GetConfig(*logger)
		return nil
	},
}
