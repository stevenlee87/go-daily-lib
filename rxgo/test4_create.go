package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

/*
传入一个[]rxgo.Producer的切片，其中rxgo.Producer的类型为func(ctx context.Context, next chan<- Item)。
我们可以在代码中调用rxgo.Of(value)生成数据，rxgo.Error(err)生成错误，然后发送到next通道中：
*/
func main() {
	observable := rxgo.Create([]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
		next <- rxgo.Of(1)
		next <- rxgo.Of(2)
		next <- rxgo.Of(3)
		next <- rxgo.Error(errors.New("unknown"))
		next <- rxgo.Of(4)
		next <- rxgo.Of(5)
	}})

	ch := observable.Observe()
	for item := range ch {
		if item.Error() {
			fmt.Println("error:", item.E)
		} else {
			fmt.Println(item.V)
		}
	}
}
