package main

import (
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

// ElementAt()只发送指定索引的数据，如ElementAt(2)只发送索引为 2 的数据，即第 3 个数据。
func main() {
	observable := rxgo.Just(0, 1, 2, 3, 4)().ElementAt(2)

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
