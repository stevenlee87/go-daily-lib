package main

import (
	"context"
	"fmt"
	"time"
)

func HandelRequest(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running, parameter: ", ctx.Value("parameter"))
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx := context.WithValue(context.Background(), "parameter", "1")
	go HandelRequest(ctx)

	time.Sleep(10 * time.Second)
}

/*
上例main()中通过WithValue()方法获得一个context，需要指定一个父context、key和value。然后通将该 context传递给子协程HandelRequest，
子协程可以读取到context的key-value。

注意：本例中子协程无法自动结束，因为context是不支持cancle的，也就是说<-ctx.Done()永远无法返回。
如果需要返回，需要在创建context时指定一个可以cancel的context作为父节点，使用父节点的cancel()在适当的时机结束整个context。
*/
