package main

import (
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

/*
FromChannel可以直接从一个已存在的<-chan rxgo.Item对象中创建 Observable：
注意：

通道需要手动调用close()关闭，上面Create()方法内部rxgo自动帮我们执行了这个步骤。
*/
func main() {
	ch := make(chan rxgo.Item)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()

	observable := rxgo.FromChannel(ch)
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
