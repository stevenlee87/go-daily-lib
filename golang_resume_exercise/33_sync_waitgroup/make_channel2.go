package main

import "fmt"

func main() {
	chanFoo := make(chan bool) // chan 没容量的时候 写和读 必须是同时的，在没有人读的情况下 就不能写入

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
