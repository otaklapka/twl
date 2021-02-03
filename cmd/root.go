package cmd

import (
	"fmt"
	"github.com/otaklapka/twl/internal"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var DbManager *internal.DbManager

var rootCmd = &cobra.Command{
	Use:   "twl <command> [<args>]",
	Short: "Tiny personal work logger",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	ex, exErr := os.Executable()
	if exErr != nil {
		fmt.Printf("Could not get executable path to create DB under: %v\n", exErr)
		os.Exit(1)
	}
	exPath := filepath.Dir(ex)

	var err error
	err, DbManager = internal.NewDbManager(fmt.Sprintf("%s/worklog.db", exPath))
	if err != nil {
		fmt.Printf("Could not init db: %v\n", err)
		os.Exit(1)
	}
}
