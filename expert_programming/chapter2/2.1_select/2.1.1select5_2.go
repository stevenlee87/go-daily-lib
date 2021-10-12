package main

import "fmt"

func SelectForChan(c chan string) {
	var recv string
	send := "Hello"

	select {
	case recv = <-c:
		fmt.Println("received %s\n", recv)
	case c <- send:
		fmt.Printf("sent %s\n", send)
	default:
		fmt.Println("all goroutines are asleep")
	}

}

func main() {
	c := make(chan string)
	SelectForChan(c)

}
