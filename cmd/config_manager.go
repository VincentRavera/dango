package cmd

import (
	"log"
	"os"

	"github.com/VincentRavera/dango/data"
	"github.com/spf13/cobra"
)

var Config = &cobra.Command{
	Use:    "config",
	Short:  "read config",
	Long:   "Ensure config is readable.",
	Args:   cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := log.Default()
		GetConfig(*logger)
		return nil
	},
}

func processErrs(e error, l log.Logger) {
	if e != nil {
		l.Fatal(e)
		os.Exit(1)
	}
}

func GetConfig(l log.Logger) data.RootConfig {
	rootpath := os.Getenv("DANGO_ROOT")
	if len(rootpath) == 0 {
		var e error
		rootpath, e = os.Getwd()
		processErrs(e, l)
	}
	l.Printf("DANGO_ROOT=%s\n", rootpath)

	batch := false
	batch_mode := os.Getenv("DANGO_BATCH")
	if len(batch_mode) > 0 {
		batch = true
	}
	l.Printf("DANGO_BATCH=%t\n", batch)
	config := data.RootConfig{
		RootPath: rootpath,
		Batch: batch,
	}
	return config
}
