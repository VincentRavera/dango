package main

import (
	"os"
	"github.com/VincentRavera/dango/cmd"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "dango",
		Short: "Manage multiple repository as one.",
		Long: "Dango is a tool to configure, build, release, multiple projects.",
	}
)

func init() {
	rootCmd.AddCommand(cmd.Config)
	rootCmd.AddCommand(cmd.Init)
	rootCmd.AddCommand(cmd.Add)
	rootCmd.AddCommand(cmd.Clone)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
