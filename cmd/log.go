package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var logCmd = &cobra.Command{
	Use:   "log <message>",
	Short: "Insert new record",
	Args: cobra.MinimumNArgs(1), // <message>
	Run: func(cmd *cobra.Command, args []string) {
		logTime, _ := cmd.Flags().GetString("time")

		pt, err := time.Parse("02.01.2006 15:04", logTime)
		if err != nil {
			fmt.Println("Time is not in \"02.01.2006 15:04\" format")
			os.Exit(1)
		}

		row, err := DbManager.Insert(args[0], pt)
		if err != nil {
			fmt.Printf("Insert error: %v\n", err)
			os.Exit(1)
		}

		id, _ := row.LastInsertId()

		fmt.Println("Added new record:")
		green := color.New(color.FgHiGreen).SprintFunc()
		black := color.New(color.FgHiBlack).SprintFunc()
		fmt.Printf("  %v  %v  %v", black(id), green(pt.Format("02.01.2006 15:04")), args[0])
	},
}

func init() {
	rootCmd.AddCommand(logCmd)

	logCmd.Flags().StringP("time", "t", time.Now().Format("02.01.2006 15:04"), "Custom time in format \"02.01.2006 15:04\"")
}
