package main

import (
	"fmt"
)

func main() {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}

/*
解析：
defer 关键字的实现跟go关键字很类似，不同的是它调⽤的是 runtime.deferproc ⽽不是 runtime.newproc 。

在 defer 出现的地⽅，插⼊了指令 call runtime.deferproc ，然后在函数返回之前的地⽅，插⼊指令
call runtime.deferreturn 。

goroutine的控制结构中，有⼀张表记录 defer ，调⽤ runtime.deferproc时会将需要defer的表达式记录在表中，
⽽在调⽤ runtime.deferreturn 的时候，则会依次从defer表中出栈并执⾏。

因此，题⽬最后输出顺序应该是 defer 定义顺序的倒序。 panic 错误并不能终⽌ defer 的执⾏。
*/
