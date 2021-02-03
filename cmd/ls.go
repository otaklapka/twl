package cmd

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"time"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List records",
	Run: func(cmd *cobra.Command, args []string) {
		listDate, _ := cmd.Flags().GetString("date")
		isListLastDate, _ := cmd.Flags().GetBool("last")

		dateToList := time.Now()

		if isListLastDate {
			if lastDate, err := DbManager.GetLastInsertDate(); err == nil {
				dateToList = lastDate
			}
		} else {
			if pt, err := time.Parse("02.01.2006", listDate); err == nil {
				dateToList = pt
			} else {
				fmt.Println("Time is not in \"02.01.2006\" format")
				os.Exit(1)
			}
		}

		var id int
		var message string
		var messageTime time.Time

		fmt.Printf("\n  Listing records of %v\n", dateToList.Format("02.01.2006"))

		table := tablewriter.NewWriter(os.Stdout)
		table.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
		table.SetCenterSeparator("")
		table.SetColumnSeparator("")
		table.SetHeader([]string{"", "", ""})
		table.SetHeaderLine(false)
		table.SetColumnColor(
			tablewriter.Colors{tablewriter.FgHiBlackColor},
			tablewriter.Colors{tablewriter.FgHiGreenColor},
			tablewriter.Colors{},
		)

		rows, _ := DbManager.List(dateToList)
		for rows.Next() {
			_ = rows.Scan(&id, &message, &messageTime)
			table.Append([]string{strconv.Itoa(id), messageTime.Format("15:04"), message})
		}

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	lsCmd.Flags().StringP("date", "d", time.Now().Format("02.01.2006"), "Custom date in format \"02.01.2006\"")
	lsCmd.Flags().BoolP("last", "l", false, "Last date with logs, not today")
}
