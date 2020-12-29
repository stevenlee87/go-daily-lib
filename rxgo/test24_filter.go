package main

import (
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

// Filter()接受一个类型为func (i interface{}) bool的参数，通过的数据使用这个函数断言，返回true的将发送给下一个阶段。否则，丢弃。
func main() {
	observable := rxgo.Range(1, 10)

	observable = observable.Filter(func(i interface{}) bool {
		return i.(int)%2 == 0
	})

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}

/*
2
4
6
8
10
*/
