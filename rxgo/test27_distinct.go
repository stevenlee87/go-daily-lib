package main

import (
	"context"
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

/*
Distinct()会记录它发送的所有数据，它不会发送重复的数据。
由于数据格式多样，Distinct()要求我们提供一个函数，根据原数据返回一个唯一标识码（有点类似哈希值）。基于这个标识码去重。

10.11.155.12, 10.22.254.128, 10.33.166.43, 10.14.222.142, 10.21.165.72: IP地址不能去重

*/
func main() {
	observable := rxgo.Just(1, 2, 2, 3, 3, 4, 4)().
		Distinct(func(_ context.Context, i interface{}) (interface{}, error) {
			return i, nil
		})
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
