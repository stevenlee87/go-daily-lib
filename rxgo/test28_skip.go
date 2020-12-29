package main

import (
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

// Skip可以跳过前若干个数据。

func main() {
	observable := rxgo.Just(1, 2, 3, 4, 5)().Skip(2)
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
