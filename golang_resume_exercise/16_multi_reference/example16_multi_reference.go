package main

import "fmt"

type query func(string) string

func exec(name string, vs ...query) string { // var vs []query
	//fmt.Println("vs --->", vs)
	ch := make(chan string)
	fn := func(i int) {
		fmt.Println("vs[i](name) --->", vs[i](name))
		ch <- vs[i](name) // vs[0]("111")
	}

	for i, _ := range vs {
		//fmt.Println(i)
		go fn(i)
	}
	return <-ch
}

func main() {
	ret := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})
	fmt.Println(ret)

}

/*
依据4个goroutine的启动后执⾏效率，很可能打印111func4，但其他的111func*也
可能先执⾏，exec只会返回⼀条信息。

上⾯的代码有严重的内存泄漏问题，出错的位置是 go fn(i) ，实际上代码执⾏后会启动4个协程，
但是因为 ch 是⾮缓冲的，只可能有⼀个协程写⼊成功。⽽其他三个协程会⼀直在后台等待写⼊。
*/
