package main

import "fmt"

func foo() {
	defer fmt.Print("A")
	defer fmt.Print("B")

	fmt.Print("C")
	panic("demopanic")
	defer fmt.Print("D")
}

func main() {
	defer func() {
		fmt.Print("demo3")
	}()

	go foo()
}

/*
Output:
demo3
*/
