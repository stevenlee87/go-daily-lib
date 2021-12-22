package main

import "fmt"

func main() {
	defer func() {
		recover()
	}()

	defer fmt.Print("A")

	defer func() {
		fmt.Print("B")
		panic("panic in defer")
		fmt.Print("C")
	}()

	panic("panic")

	fmt.Print("D")
}
