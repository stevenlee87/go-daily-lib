package main

import (
	"fmt"
	"sync"
)

var strPool = sync.Pool{
	New: func() interface{} {
		return "test str"
	},
}

func main() {
	str := strPool.Get()
	fmt.Println(str)
	strPool.Put(str)
}

/*
1.通过New去定义你这个池子里面放的究竟是什么东西，在这个池子里面你只能放一种类型的东西。比如在上面的例子中我就在池子里面放了字符串。
2.我们随时可以通过Get方法从池子里面获取我们之前在New里面定义类型的数据。
3.当我们用完了之后可以通过Put方法放回去，或者放别的同类型的数据进去。

*/
