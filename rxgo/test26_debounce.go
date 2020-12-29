package main

import (
	"fmt"
	"time"

	"github.com/reactivex/rxgo/v2"
)

// Debounce()比较有意思，它收到数据后还会等待指定的时间间隔，后续间隔内没有收到其他数据才会发送刚开始的数据。

func main() {
	ch := make(chan rxgo.Item)

	go func() {
		ch <- rxgo.Of(1)
		time.Sleep(2 * time.Second)
		ch <- rxgo.Of(2)
		ch <- rxgo.Of(3)
		time.Sleep(2 * time.Second)
		close(ch)
	}()

	observable := rxgo.FromChannel(ch).Debounce(rxgo.WithDuration(1 * time.Second))
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}

/*
1
3
*/
