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
	go func(ch chan int) {
		time.Sleep(time.Second)
		<-ch
	}(ch)
	c := time.Tick(1 * time.Second)
	for range c {
		// 使用 runtime.NumGoroutine 检查进程创建的 goroutine 数量总数
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine()) // main 是主协程
	}
}

/*
⼀段时间后总是输出 #goroutines: 2

解析：
默认情况下 make 初始化的 channel 是⽆缓冲的，也就是在迭代写时会阻塞。

有缓冲区channel
channel 的缓冲区为1，向channel发送第一个数据，主协程不会退出。
发送第二个时候，缓冲区已经满了，此时阻塞主协程。
*/
