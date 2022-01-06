package main

import (
	"fmt"
	"time"
)

func addNumberToChan(chanName chan int) {
	for {
		chanName <- 1
		time.Sleep(1 * time.Second)
	}
}

/* 通过这个例子可以看出，select的case语句读管道时不会阻塞，尽管管道中没有数据。
这是由于case语句编译后调用读管道时会明确传入不阻塞的参数，读不到数据时不会讲当前协程加入等待队列，
而是直接返回
*/
func main() {
	var chan1 = make(chan int, 10)
	var chan2 = make(chan int, 10)

	go addNumberToChan(chan1)
	go addNumberToChan(chan2)

	for {
		select {
		case e := <-chan1:
			fmt.Printf("Get element from chan1: %d\n", e)
		case e := <-chan2:
			fmt.Printf("Get element from chan2: %d\n", e)
		default:
			fmt.Printf("No element in chan1 and chan2.\n")
			time.Sleep(1 * time.Second)
		}
	}
}
