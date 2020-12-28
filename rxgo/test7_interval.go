package main

import (
	"fmt"
	"time"

	"github.com/reactivex/rxgo/v2"
)

/*
Interval以传入的时间间隔生成一个无穷的数字序列，从 0 开始：
*/
func main() {
	observable := rxgo.Interval(rxgo.WithDuration(5 * time.Second))
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
