package main

import (
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

/*
除了使用for range之外，我们还可以调用 Observable 的ForEach()方法来实现遍历。ForEach()接受 3 个回调函数：

NextFunc：类型为func (v interface {})，处理数据；
ErrFunc：类型为func (err error)，处理错误；
CompletedFunc：类型为func ()，Observable 完成时调用。
有点Promise那味了。使用ForEach()，可以将上面的示例改写为：
*/

func main() {
	observable := rxgo.Just(1, 2, 3, 4, 5)()
	//observable.ForEach(func(v interface{}) {
	<-observable.ForEach(func(v interface{}) {
		fmt.Println("received:", v)
	}, func(err error) {
		fmt.Println("error:", err)
	}, func() {
		fmt.Println("completed")
	})
}
