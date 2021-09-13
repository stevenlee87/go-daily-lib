package main

func main() {
	c := make(chan int, 3)
	close(c)
	c <- 1
}

/*
panic: send on closed channel

goroutine 1 [running]:
main.main()
*/
