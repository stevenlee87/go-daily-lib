package main

import "fmt"

type query func(string) string

func exec(name string, vs ...query) string { // var vs []query fslice := []int(arg)
	ch := make(chan string)
	fn := func(i int) {
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
*/
