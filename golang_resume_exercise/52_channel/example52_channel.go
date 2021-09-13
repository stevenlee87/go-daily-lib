package main

import "fmt"

func main() {
	fmt.Println("以下是数值的chan")
	ci := make(chan int, 3)
	ci <- 1
	close(ci)
	num, ok := <-ci
	fmt.Printf("读chan的协程结束，num=%v， ok=%v\n", num, ok)
	num1, ok1 := <-ci
	fmt.Printf("再读chan的协程结束，num=%v， ok=%v\n", num1, ok1)
	num2, ok2 := <-ci
	fmt.Printf("再再读chan的协程结束，num=%v， ok=%v\n", num2, ok2)

	fmt.Println("以下是字符串chan")
	cs := make(chan string, 3)
	cs <- "aaa"
	close(cs)
	str, ok := <-cs
	fmt.Printf("读chan的协程结束，str=%v， ok=%v\n", str, ok)
	str1, ok1 := <-cs
	fmt.Printf("再读chan的协程结束，str=%v， ok=%v\n", str1, ok1)
	str2, ok2 := <-cs
	fmt.Printf("再再读chan的协程结束，str=%v， ok=%v\n", str2, ok2)

	fmt.Println("以下是结构体chan")
	type MyStruct struct {
		Name string
	}
	cstruct := make(chan MyStruct, 3)
	cstruct <- MyStruct{Name: "haha"}
	close(cstruct)
	stru, ok := <-cstruct
	fmt.Printf("读chan的协程结束，stru=%v， ok=%v\n", stru, ok)
	stru1, ok1 := <-cs
	fmt.Printf("再读chan的协程结束，stru=%v， ok=%v\n", stru1, ok1)
	stru2, ok2 := <-cs
	fmt.Printf("再再读chan的协程结束，stru=%v， ok=%v\n", stru2, ok2)
}

/*
读已经关闭的 chan 能⼀直读到东⻄，但是读到的内容根据通道内关闭前是否有元素⽽不同。
如果 chan 关闭前，buffer 内有元素还未读 , 会正确读到 chan 内的值，且返回的第⼆ 个 bool 值（是否读成功）为 true。
如果 chan 关闭前，buffer 内有元素已经被读完，chan 内⽆值，接下来所有接收的值都会⾮阻塞直接成功，返回 channel 元素的零值，但是第⼆个 bool 值⼀直为 false。

写已经关闭的 chan 会 panic
*/
