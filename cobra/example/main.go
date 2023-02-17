package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	// Create the command object.
	cmd := &cobra.Command{
		Use:   "example",
		Short: "A brief description of the command",
		Long:  `A longer description of the command that spans multiple lines.`,
		Run:   run,
	}

	// Add flags and configuration settings.
	cmd.Flags().IntP("num", "n", 0, "The number to use in the command")

	// Execute the command.
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func run(cmd *cobra.Command, args []string) {
	// Get the value of the -n flag.
	num, _ := cmd.Flags().GetInt("num")

	// Do something with the value.
	fmt.Println("The number is", num)
}

/*
./main -n 1
The number is 1
*/
