package main

import "fmt"

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := Fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Printf("Fibonacci: %d\n", f())
	}
}

/*
go build -gcflags=-m
Fibonacci()函数中原本属于局部变量的a和b由于闭包的引用，不得不将二者放到堆中，以致产生逃逸：
./4.3.2fibonacci.go:7:9: can inline Fibonacci.func1
./4.3.2fibonacci.go:17:13: inlining call to fmt.Printf
./4.3.2fibonacci.go:6:2: moved to heap: a
./4.3.2fibonacci.go:6:5: moved to heap: b
./4.3.2fibonacci.go:7:9: func literal escapes to heap
./4.3.2fibonacci.go:17:34: f() escapes to heap
./4.3.2fibonacci.go:17:13: []interface {}{...} does not escape
<autogenerated>:1: .this does not escape

栈上分配内存比堆中分配内存有更高的效率
栈上分配内存不需要GC处理
堆上分配的内存使用完毕会交给GC处理
逃逸分析的目的是决定分配地址是栈还是堆
逃逸分析在编译阶段完成

思考一下这个问题：函数传递指针真的比传值的效率高吗？
我们知道传递指针可以减少底层值的复制，可以提高效率，但是如果复制的数据量小，由于指针传递会产生逃逸，则可能使用堆，也可能增加GC的负担，
所以传递指针不一定是高效的。
*/
