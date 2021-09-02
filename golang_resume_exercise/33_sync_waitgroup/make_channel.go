package main

import "fmt"

func main() {
	chanFoo := make(chan bool, 1) // the only difference is the buffer size of 1

	for i := 0; i < 5; i++ {
		select {
		case <-chanFoo:
			fmt.Println("Read")
		case chanFoo <- true:
			fmt.Println("Write")
		default:
			fmt.Println("Neither")
		}
	}
}
