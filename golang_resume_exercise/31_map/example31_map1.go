package main

/*
⾼并发下的锁与map的读写
场景：在⼀个⾼并发的web服务器中，要限制IP的频繁访问。现模拟100个IP同时并发访问服 务器，每个IP要重复访问1000次。
每个IP三分钟之内只能访问⼀次。修改以下代码完成该过程，要求能成功输出 success:100
*/

import (
	"fmt"
	"time"
)

type Ban struct {
	visitIPs map[string]time.Time
}

func NewBan() *Ban {
	return &Ban{visitIPs: make(map[string]time.Time)}
}
func (o *Ban) visit(ip string) bool {
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = time.Now()
	return false
}
func main() {
	success := 0
	ban := NewBan()
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func() {
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					success++
				}
			}()
		}

	}
	fmt.Println("success:", success)
}

/*
解析该问题主要考察了并发情况下map的读写问题，⽽给出的初始代码，⼜存在 for 循环中启动 goroutine 时变量使⽤问题
以及 goroutine 执⾏滞后问题。

因此，⾸先要保证启动的 goroutine 得到的参数是正确的，然后保证 map 的并发读写，最后保证三分钟只能访问⼀次。
多CPU核⼼下修改 int 的值极端情况下会存在不同步情况，因此需要原⼦性的修改int值。
下⾯给出的实例代码，是启动了⼀个协程每分钟检查⼀下 map 中的过期 ip ， for 启动协程时传参。
*/
