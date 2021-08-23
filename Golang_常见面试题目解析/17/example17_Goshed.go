package main

import (
	"fmt"
	"runtime"
)

func main() {
	var i byte
	go func() {
		for i = 0; i <= 255; i++ {
		}
	}()
	fmt.Println("Dropping mic")
	// Yield execution to force executing other goroutines
	runtime.Gosched()
	runtime.GC()
	fmt.Println("Done")
}

/*
解析：
Golang 中，byte 其实被 alias 到 uint8 上了。所以上⾯的 for 循环会始终成⽴，因为
i++ 到 i=255 的时候会溢出，i <= 255 ⼀定成⽴。
也即是， for 循环永远⽆法退出，所以上⾯的代码其实可以等价于这样：
正在被执⾏的 goroutine 发⽣以下情况时让出当前 goroutine 的执⾏权，并调度后⾯的
goroutine 执⾏：
IO 操作
Channel 阻塞
system call
运⾏较⻓时间

如果⼀个 goroutine 执⾏时间太⻓，scheduler 会在其 G 对象上打上⼀个标志（preempt），
当这个 goroutine 内部发⽣函数调⽤的时候，会先主动检查这个标志，如果为 true 则会让出执⾏权。

main 函数⾥启动的 goroutine 其实是⼀个没有 IO 阻塞、没有 Channel 阻塞、没有system call、
没有函数调⽤的死循环。也就是，它⽆法主动让出⾃⼰的执⾏权，即使已经执⾏很⻓时间，scheduler 已经标志
了 preempt。

⽽ golang 的 GC 动作是需要所有正在运⾏ goroutine 都停⽌后进⾏的。因此，程序
会卡在 runtime.GC() 等待所有协程退出。
*/
