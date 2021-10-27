package main

import (
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		for {
			// 无函数调用的无限循环
		}
	}()
	time.Sleep(1 * time.Second) // 系统调用，出让执行权给上面的协程
	println("Done")
}

/*
上面的代码在Go 1.14之前会陷入协程的无限循环中，协程永远无法被抢占，导致主协程无法继续执行。
直到在Go1.14中，调度器引入了基于信号的抢占机制，这个问题才得到了解决
*/
