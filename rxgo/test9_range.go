package main

import (
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

// Range可以生成一个范围内的数字：
func main() {
	observable := rxgo.Range(0, 3)
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
