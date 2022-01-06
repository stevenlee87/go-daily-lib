package main

import (
	"fmt"
	"time"
)

/*
有时我们希望从一个管道中读取数据，在管道中没有数据时，我们不想让程序永远阻塞在管道中，
而是设定一个超时时间，在此时间段中如果管道中还是没有数据到来，则判定为超时。

Go源码包中有大量类似的用法，比如从一个连接中等待数据，其简单的用法如下代码所示：
*/
func WaitChannel(conn <-chan string) bool {
	timer := time.NewTimer(1 * time.Second)

	select {
	case <-conn:
		timer.Stop()
		fmt.Println("timer stop")
		return true
	case <-timer.C: // 超时
		println("WaitChannel timeout!")
		return false
	}
}

func main() {
	var conn <-chan string
	WaitChannel(conn)
}

/*
WaitChannel作用就是检测指定的管道中是否有数据到来，通过select语句轮询conn和timer.C两个管道，
timer 会在1s后向timer.C写入数据，如果1s内conn还没有数据，则会判断为超时。
*/
