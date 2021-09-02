package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
⾼并发下的锁与map的读写
场景：在⼀个⾼并发的web服务器中，要限制IP的频繁访问。现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。
每个IP三分钟（按照代码看是一分钟？？？）之内只能访问⼀次。修改以下代码完成该过程，要求能成功输出 success:100

解析该问题主要考察了并发情况下map的读写问题，⽽给出的初始代码，⼜存在 for 循环中启动 goroutine 时变量使⽤问题
以及 goroutine 执⾏滞后问题。

因此，⾸先要保证启动的 goroutine 得到的参数是正确的，然后保证 map 的并发读写，最后保证三分钟只能访问⼀次。
多CPU核⼼下修改 int 的值极端情况下会存在不同步情况，因此需要原⼦性的修改int值。
下⾯给出的实例代码，是启动了⼀个协程每分钟检查⼀下 map 中的过期 ip ， for 启动协程时传参。
*/

type Ban struct {
	visitIPs map[string]time.Time
	lock     sync.Mutex
}

func NewBan(ctx context.Context) *Ban {
	o := &Ban{visitIPs: make(map[string]time.Time)}
	go func() {
		timer := time.NewTimer(time.Minute * 1)
		for {
			select {
			case <-timer.C:
				o.lock.Lock()
				for k, v := range o.visitIPs {
					if time.Now().Sub(v) >= time.Minute*1 {
						delete(o.visitIPs, k)
					}
				}
				o.lock.Unlock()
				timer.Reset(time.Minute * 1)
			case <-ctx.Done():
				return
			}
		}
	}()
	return o
}

func (o *Ban) visit(ip string) bool {
	o.lock.Lock()
	defer o.lock.Unlock()
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = time.Now()
	//fmt.Println(ip)
	return false
}

func main() {
	success := int64(0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ban := NewBan(ctx)

	wait := &sync.WaitGroup{}

	wait.Add(1000 * 100)
	for i := 0; i < 1000; i++ {
		//fmt.Println(i)
		for j := 0; j < 100; j++ {
			go func(j int) {
				defer wait.Done()
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					atomic.AddInt64(&success, 1)
				}
			}(j)
		}
	}
	wait.Wait()
	fmt.Println("success:", success)
}
