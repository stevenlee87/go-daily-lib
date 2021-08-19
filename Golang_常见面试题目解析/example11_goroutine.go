package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}

/*
解析：
在 golang 中 goroutine 的调度时间是不确定的，
在题⽬中，第⼀个写channel 的 goroutine 可能还未调⽤，或已调⽤但没有写完时直接 close 管道，
可能导致写失败，即出现 panic 错误。
*/
