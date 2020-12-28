package main

import (
	"context"
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

/*
可以给Start方法传入[]rxgo.Supplier作为参数，它可以包含任意数量的rxgo.Supplier类型。rxgo.Supplier的底层类型为：

// rxgo/types.go
var Supplier func(ctx context.Context) rxgo.Item
Observable 内部会依次调用这些rxgo.Supplier生成rxgo.Item：

*/
func Supplier1(ctx context.Context) rxgo.Item {
	return rxgo.Of(1)
}

func Supplier2(ctx context.Context) rxgo.Item {
	return rxgo.Of(2)
}

func Supplier3(ctx context.Context) rxgo.Item {
	return rxgo.Of(3)
}

func main() {
	observable := rxgo.Start([]rxgo.Supplier{Supplier1, Supplier2, Supplier3})
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
