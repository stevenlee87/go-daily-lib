package main

import "fmt"

func main() {
	i, j := 0, 0
	if true {
		j, k := 1, 1
		fmt.Printf("j = %d, k = %d\n", j, k)
	}

	fmt.Printf("i = %d, j = %d", i, j)
}
