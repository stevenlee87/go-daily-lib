package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("A")
		}
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("B")
		}
	}()

	panic("demo")
	fmt.Println("C")
}

/*
Output:
B
*/
