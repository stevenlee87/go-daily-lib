package main

import "fmt"

func main() {
	a := 7
	remainder_sum := a % 10
	division_sum := a / 10
	fmt.Printf("remainder_sum is: %d, division_sum is: %d\n", remainder_sum, division_sum)
	// 7 0

	b := 10
	remainder_sum = b % 10
	division_sum = b / 10
	fmt.Printf("remainder_sum is: %d, division_sum is: %d\n", remainder_sum, division_sum)
	// 0 1

	c := 8
	remainder_sum = c % 10
	division_sum = c / 10
	fmt.Printf("remainder_sum is: %d, division_sum is: %d\n", remainder_sum, division_sum)
	// 8 0
}
