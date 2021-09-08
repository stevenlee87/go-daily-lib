package main

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var pool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

func processRequest(size int) {
	b := pool.Get().(*bytes.Buffer)
	time.Sleep(500 * time.Millisecond)
	b.Grow(size)
	pool.Put(b)
	time.Sleep(1 * time.Millisecond)
}

func main() {
	go func() {
		for {
			processRequest(1 << 28) // 256MiB
		}
	}()
	for i := 0; i < 1000; i++ {
		//for i := 0; i < 10; i++ {
		go func() {
			for {
				processRequest(1 << 10) // 1KiB
			}
		}()
	}
	var stats runtime.MemStats
	for i := 0; ; i++ {
		runtime.ReadMemStats(&stats)
		fmt.Printf("Cycle %d: %dB\n", i, stats.Alloc)
		time.Sleep(time.Second)
		runtime.GC()
	}
}

/*
可以编译，运⾏时内存先暴涨，但是过⼀会会回收掉.
1000次循环有点儿大，物理内存使用了30G都没有回收掉。
如果把for i := 0; i < 10; i++ {循环改成10的话，内存在循环100次左右就回收掉一部分内存.

sync.Pool 是并发安全的吗？
sync.Pool 当然是并发安全的。官方文档里明确说了：
A Pool is safe for use by multiple goroutines simultaneously.
但是，为什么我这里会单独提出来呢？
因为 sync.Pool 只是本身的 Pool 数据结构是并发安全的，并不是说 Pool.New 函数一定是线程安全的。
Pool.New 函数可能会被并发调用 ，如果 New 函数里面的实现是非并发安全的，那就会有问题。

*/
