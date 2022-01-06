package main

import "fmt"

func main() {
	err := 1
	i := 0
	for err != 123 {
		if i == 2 {
			err = 123
		}

		fmt.Println(i)
		i = i + 1
	}
}
