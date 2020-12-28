package main

import (
	"context"
	"fmt"
	"time"

	"github.com/reactivex/rxgo/v2"
)

// 我们通过在创建 Observable 的方法中指定rxgo.WithPublishStrategy()选项就可以创建可连接的 Observable：
func main() {
	ch := make(chan rxgo.Item)
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()

	observable := rxgo.FromChannel(ch, rxgo.WithPublishStrategy())

	observable.DoOnNext(func(i interface{}) {
		fmt.Printf("First observer: %d\n", i)
	})

	time.Sleep(3 * time.Second)
	fmt.Println("before subscribe second observer")

	observable.DoOnNext(func(i interface{}) {
		fmt.Printf("Second observer: %d\n", i)
	})

	observable.Connect(context.Background())
	time.Sleep(3 * time.Second)
	fmt.Print("end!")
}

/*
before subscribe second observer
Second observer: 1
First observer: 1
First observer: 2
First observer: 3
Second observer: 2
Second observer: 3
end!

上面是等两个观察者都注册之后，并且手动调用了 Observable 的Connect()方法才产生数据。
而且可连接的 Observable 有一个特性：它是冷启动的！！！，即每个观察者都会收到一份相同的拷贝。
*/
