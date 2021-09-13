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

//在 src/runtime/chan.go func chansend(c *hchan,ep unsafe.Pointer,block bool,callerpc uintptr) bool
{

	//省略其他
	if c.closed != 0 {
		unlock(&c.lock)
		panic(plainError("send on closed channel"))
	}
	//省略其他
}

*/
