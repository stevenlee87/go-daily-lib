package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
	//time.Sleep(time.Second * 1)
}

/* golang 1.16
hello
hello

world
hello
world
hello

world
world
hello
hello


*/
