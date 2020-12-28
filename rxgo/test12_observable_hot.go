package main

import (
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

/*
根据数据在何处生成，Observable 被分为 Hot 和 Cold 两种类型（类比热启动和冷启动）。
数据在其它地方生成的被成为 Hot Observable。相反，在 Observable 内部生成数据的就是 Cold Observable。
使用上面介绍的方法创建的实际上都是 Hot Observable。

*/
func main() {
	ch := make(chan rxgo.Item)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()

	observable := rxgo.FromChannel(ch)

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}

/*
上面创建的是 Hot Observable。但是有个问题，第一次Observe()消耗了所有的数据，第二个就没有数据输出了。

而 Cold Observable 就不会有这个问题，因为它创建的流是独立于每个观察者的。即每次调用Observe()都创建一个新的 channel。
我们使用Defer()方法创建 Cold Observable，它的参数与Create()方法一样。
*/
