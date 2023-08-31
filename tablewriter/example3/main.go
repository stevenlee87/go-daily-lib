package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

// Example 3 - CSV
func main() {
	fmt.Println(os.Getwd())
	table, _ := tablewriter.NewCSV(os.Stdout, "./example3/test_info.csv", true)
	table.SetAlignment(tablewriter.ALIGN_LEFT) // Set Alignment
	table.Render()
}
