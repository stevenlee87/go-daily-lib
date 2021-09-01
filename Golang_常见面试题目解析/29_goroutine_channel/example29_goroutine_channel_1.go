package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
写代码实现两个 goroutine，其中⼀个产⽣随机数并写⼊到 go channel 中，另外⼀
个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。
解析
这是⼀道很简单的golang基础题⽬，实现⽅法也有很多种，⼀般想答让⾯试官满意的答案还是有⼏点注意的地⽅。
1. goroutine 在golang中是⾮阻塞的
2. channel ⽆缓冲情况下，读写都是阻塞的，且可以⽤ for 循环来读取数据，当管道关闭后， for 退出。
3. golang 中有专⽤的 select case 语法从管道读取数据。
示例代码如下：
*/

func main() {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			out <- rand.Intn(5)
		}
		close(out)
	}()
	go func() {
		defer wg.Done()
		for i := range out {
			fmt.Println(i)
		}
	}()
	wg.Wait()
}
