package main

import (
	"fmt"
	"time"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	ch := make(chan rxgo.Item)
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()

	observable := rxgo.FromChannel(ch)

	observable.DoOnNext(func(i interface{}) {
		fmt.Printf("First observer: %d\n", i)
	})

	time.Sleep(3 * time.Second)
	fmt.Println("before subscribe second observer")

	observable.DoOnNext(func(i interface{}) {
		fmt.Printf("Second observer: %d\n", i)
	})

	time.Sleep(3 * time.Second)
}

/*
上例中我们使用DoOnNext()方法来注册观察者。
由于DoOnNext()方法是异步执行的，所以为了等待结果输出，在最后增加了一行time.Sleep。运行：

$ go run main.go
First observer: 1
First observer: 2
First observer: 3
before subscribe second observer
由输出可以看出，注册第一个观察者之后就开始产生数据了。
*/
