package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var ch chan int
	go func() {
		ch = make(chan int, 1)
		ch <- 1
	}()
	go func(ch *chan int) {
		time.Sleep(time.Second)
		res := <-*ch
		fmt.Println(res)
	}(&ch)
	fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	c := time.Tick(1 * time.Second)
	for range c {
		// 使用 runtime.NumGoroutine 检查进程创建的 goroutine 数量总数
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}
