package main

import (
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	count := 3
	observable := rxgo.Range(0, 10).GroupBy(count, func(item rxgo.Item) int {
		return item.V.(int) % count
	}, rxgo.WithBufferedChannel(10))

	for subObservable := range observable.Observe() {
		fmt.Println("New observable:")

		for item := range subObservable.V.(rxgo.Observable).Observe() {
			fmt.Printf("item: %v\n", item.V)
		}
	}
}

/*
上面根据每个数模 3 的余数将整个流分为 3 组。运行：

$ go run main.go
New observable:
item: 0
item: 3
item: 6
item: 9
New observable:
item: 1
item: 4
item: 7
item: 10
New observable:
item: 2
item: 5
item: 8
注意rxgo.WithBufferedChannel(10)的使用，由于我们的数字是连续生成的，依次为 0->1->2->...->9->10。
而 Observable 默认是惰性的，即由Observe()驱动。
内层的Observe()在返回一个 0 之后就等待下一个数，但是下一个数 1 不在此 Observable 中。所以会陷入死锁。
使用rxgo.WithBufferedChannel(10)，设置它们之间的连接 channel 缓冲区大小为 10，
这样即使我们未取出 channel 里面的数字，上游还是能发送数字进来。

*/
