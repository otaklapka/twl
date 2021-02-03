package cmd

import (
	"fmt"
	"github.com/otaklapka/twl/internal"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"time"
)

var tdiffCmd = &cobra.Command{
	Use:   "tdiff <id or hh:mm> <id or hh:mm>",
	Short: "Count duration between records or time",
	Args: cobra.MinimumNArgs(2), // <value> <value>
	Run: func(cmd *cobra.Command, args []string) {
		time1, err := parseArg(args[0])
		if err != nil {
			fmt.Printf("First argument error: %v\n", err)
			os.Exit(1)
		}

		time2, err := parseArg(args[1])
		if err != nil {
			fmt.Printf("Second argument error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(time2.Sub(time1))
	},
}

func init() {
	rootCmd.AddCommand(tdiffCmd)
}

func parseArg(arg string) (time.Time, error) {
	t := time.Now()
	var err error
	var row *internal.LogRecord
	var id int

	if id, err = strconv.Atoi(arg); err == nil {
		if row, err = DbManager.GetRecord(id); err == nil {
			t = row.Time
		}
	} else {
		var pt time.Time
		if pt, err = time.Parse("15:04", arg); err == nil {
			t = time.Date(t.Year(), t.Month(), t.Day(), pt.Hour(), pt.Minute(), pt.Second(), 0, t.Location())
		}
	}

	return t, err
}
