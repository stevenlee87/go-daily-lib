package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 用来统计实例真正创建的次数
var numCalcsCreated int32

// 创建实例的函数
func createBuffer() interface{} {
	// 这里要注意下，非常重要的一点。这里必须使用原子加，不然有并发问题；
	atomic.AddInt32(&numCalcsCreated, 1)
	buffer := make([]byte, 1024)
	return &buffer
}

func main() {
	// 创建实例
	bufferPool := &sync.Pool{
		New: createBuffer,
	}

	// 多 goroutine 并发测试
	numWorkers := 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			// 申请一个 buffer 实例
			buffer := bufferPool.Get()
			_ = buffer.(*[]byte)
			// 释放一个 buffer 实例
			defer bufferPool.Put(buffer)
		}()
	}
	wg.Wait()
	fmt.Printf("%d buffer objects were created.\n", numCalcsCreated)
}

// 8 buffer objects were created.
/*
一、为什么 sync.Pool 不适合用于像 socket 长连接或数据库连接池?
因为，我们不能对 sync.Pool 中保存的元素做任何假设，以下事情是都可以发生的：
1. Pool 池里的元素随时可能释放掉，释放策略完全由 runtime 内部管理；
2. Get 获取到的元素对象可能是刚创建的，也可能是之前创建好 cache 住的。使用者无法区分；
3. Pool 池里面的元素个数你无法知道；
所以，只有的你的场景满足以上的假定，才能正确的使用 Pool 。sync.Pool 本质用途是增加临时对象的重用率，减少 GC 负担。划重点：临时对象。所以说，像 socket 这种带状态的，长期有效的资源是不适合 Pool 的。

二、总结
1. sync.Pool 本质用途是增加临时对象的重用率，减少 GC 负担；
2. 不能对 Pool.Get 出来的对象做预判，有可能是新的（新分配的），有可能是旧的（之前人用过，然后 Put 进去的）；
3. 不能对 Pool 池里的元素个数做假定，你不能够；
4. sync.Pool 本身的 Get, Put 调用是并发安全的，sync.New 指向的初始化函数会并发调用，里面安不安全只有自己知道；
5. 当用完一个从 Pool 取出的实例时候，一定要记得调用 Put，否则 Pool 无法复用这个实例，通常这个用 defer 完成；
*/
