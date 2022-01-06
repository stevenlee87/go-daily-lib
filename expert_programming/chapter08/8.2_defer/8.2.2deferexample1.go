package main

import "fmt"

func foo() (ret int) {
	defer func() {
		ret++
	}()

	return 0
}

func main() {
	fmt.Println(foo())
}
