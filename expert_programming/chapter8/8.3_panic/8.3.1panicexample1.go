package main

import "fmt"

func foo() {
	defer fmt.Print("A")
	defer fmt.Print("B")

	fmt.Print("C")
	panic("demo")
	defer fmt.Print("D")
}

func main() {
	defer func() {
		recover()
	}()

	foo()
}

/*
Output:
CBA
*/
