package main

import (
	"fmt"
	"time"

	"github.com/reactivex/rxgo/v2"
)

// 在已存在的 Observable 对象上调用Repeat，可以实现每隔指定时间，重复一次该序列，一共重复指定次数：
func main() {
	observable := rxgo.Just(1, 2, 3)().Repeat(
		3, rxgo.WithDuration(2*time.Second),
	)
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
