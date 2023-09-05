package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// Example 4 - Custom Separator
func main() {
	table, _ := tablewriter.NewCSV(os.Stdout, "./example4/test_info.csv", true)
	table.SetRowLine(true) // Enable row line

	// Change table lines
	table.SetCenterSeparator("*")
	table.SetColumnSeparator("╪")
	table.SetRowSeparator("-")

	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render()
}
