package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("A")
		}
	}()

	panic("demo")
	fmt.Println("B")
}

/*
Output:
A
*/
