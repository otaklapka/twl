package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"time"
)

var setCmd = &cobra.Command{
	Use:   "set <id>",
	Short: "Update record",
	Args: cobra.MinimumNArgs(1), // <id>
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Id is not a number")
			os.Exit(1)
		}

		newTime, _ := cmd.Flags().GetString("time")
		pt, err := time.Parse("02.01.2006 15:04", newTime)
		if err != nil {
			fmt.Println("Time is not in \"02.01.2006 15:04\" format")
			os.Exit(1)
		}

		newMessage, _ := cmd.Flags().GetString("message")

		if err = DbManager.Set(id, newMessage, pt); err != nil {
			fmt.Printf("Uprade error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Record updated")
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	setCmd.Flags().StringP("time", "t", time.Time{}.Format("02.01.2006 15:04"), "New log time in format \"02.01.2006 15:04\"")
	setCmd.Flags().StringP("message", "m", "", "New message")
}
