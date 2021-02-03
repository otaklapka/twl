package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm <id>",
	Short: "Remove record by id",
	Args: cobra.MinimumNArgs(1), // <id>
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Id is not a number")
			os.Exit(1)
		}

		if err = DbManager.Delete(id); err != nil {
			fmt.Printf("Delete error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Record removed")
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
