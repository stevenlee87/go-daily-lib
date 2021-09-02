package main

import (
	"fmt"
	"time"
)

/*
写出以下逻辑，要求每秒钟调⽤⼀次proc并保证程序不退出?

解析题⽬主要考察了两个知识点：
1. 定时执⾏任务
2. 捕获 panic 错误 题⽬中要求每秒钟执⾏⼀次，⾸先想到的就是 time.Ticker 对象，
该函数可每秒钟往 chan 中放⼀个 Time ,正好符合我们的要求。

在 golang 中捕获 panic ⼀般会⽤到 recover() 函数。
*/

func proc() {
	panic("ok")
}

func main() {
	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出

		t := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-t.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()
					proc()
				}()
			}
		}
	}()
	select {} // 让程序一直执行下去，没有空select。主程序就退出了。
}
