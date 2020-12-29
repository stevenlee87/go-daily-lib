package main

import (
	"fmt"
	"time"

	"github.com/reactivex/rxgo/v2"
)

/*
Buffer
Buffer按照一定的规则收集接收到的数据，然后一次性发送出去（作为切片），而不是收到一个发送一个。有 3 种类型的Buffer：

BufferWithCount(n)：每收到n个数据发送一次，最后一次可能少于n个；
BufferWithTime(n)：发送在一个时间间隔n内收到的数据；
BufferWithTimeOrCount(d, n)：收到n个数据，或经过d时间间隔，发送当前收到的数据。

*/
func main() {
	ch := make(chan rxgo.Item, 1)

	go func() {
		i := 0
		for range time.Tick(time.Second) {
			ch <- rxgo.Of(i)
			i++
		}
	}()

	observable := rxgo.FromChannel(ch).BufferWithTimeOrCount(rxgo.WithDuration(3*time.Second), 2)

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
