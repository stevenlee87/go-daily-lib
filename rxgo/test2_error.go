package main

import (
	"errors"
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

/*
实际上rxgo.Item还可以包含错误。所以在使用时，我们应该做一层判断：
我们使用item.Error()检查是否出现错误。然后使用item.V访问数据，item.E访问错误。
*/
func main() {
	observable := rxgo.Just(1, 2, errors.New("unknown"), 3, 4, 5)()
	ch := observable.Observe()
	for item := range ch {
		if item.Error() {
			fmt.Println("error:", item.E)
		} else {
			fmt.Println(item.V)
		}
	}
}
