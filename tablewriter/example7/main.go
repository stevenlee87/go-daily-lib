package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// Example 7 - Identical cells merging (specify the column index to merge)
func main() {
	data := [][]string{
		[]string{"1/1/2014", "Domain name", "1234", "$10.98"},
		[]string{"1/1/2014", "January Hosting", "1234", "$10.98"},
		[]string{"1/4/2014", "February Hosting", "3456", "$51.00"},
		[]string{"1/4/2014", "February Extra Bandwidth", "4567", "$30.00"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})
	table.SetFooter([]string{"", "", "Total", "$146.93"})
	table.SetAutoMergeCellsByColumnIndex([]int{0, 2, 3}) // 合并第0列(DATE),第2列（CV2）和第3列（AMOUNT）相同的cell
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()
}
