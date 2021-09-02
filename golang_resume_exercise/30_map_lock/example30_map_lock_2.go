package main

import (
	"fmt"
	"sync"
	"time"
)

type Map struct {
	c   map[string]*entry
	rmx *sync.RWMutex
}
type entry struct {
	ch      chan struct{}
	value   interface{}
	isExist bool
}

func (m *Map) Out(key string, val interface{}) {
	m.rmx.Lock()
	defer m.rmx.Unlock()
	if e, ok := m.c[key]; ok {
		e.value = val
		e.isExist = true
		fmt.Println(val)
	} else {
		e = &entry{ch: make(chan struct{}), isExist: true, value: val}
		m.c[key] = e
		fmt.Println(e)
		close(e.ch)
	}
}

func (m *Map) Rd(key string, timeout time.Duration) interface{} {
	m.rmx.Lock()
	if e, ok := m.c[key]; ok && e.isExist {
		fmt.Println("e.isExist", e.isExist)
		m.rmx.Unlock()
		fmt.Println("e.value", e.value)
		return e.value
	} else if !ok {
		fmt.Println("ok ???", ok)
		e = &entry{ch: make(chan struct{}), isExist: false}
		m.c[key] = e
		m.rmx.Unlock() // Unlock为什么在读取map之后执行？
		fmt.Println("协程阻塞1->", key)
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			fmt.Println("协程超时1->", key)
			return nil
		}
	} else {
		m.rmx.Unlock()
		fmt.Println("协程阻塞2->", key)
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			fmt.Println("协程超时2->", key)
			return nil
		}
	}
}

/*
实现阻塞读且并发安全的map GO⾥⾯MAP如何实现key不存在 get操作等待 直到key存在或者超时，保证并发安全，
且需要实现以下接⼝：
type sp interface {
	Out(key string, val interface{}) //存入key /val，如果该key读取的 goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回

	Rd(key string, timeout time.Duration) interface{} //读取一个key，如果 key不存在阻塞，等待key存在或者超时
}
解析：
看到阻塞协程第⼀个想到的就是 channel ，题⽬中要求并发安全，那么必须⽤锁，还要实现多个goroutine读的时候如果值不存在则阻塞，
直到写⼊值，那么每个键值需要有⼀个阻塞 goroutine 的 channel 。
实现如下：
*/
func main() {
	m := Map{
		c:   map[string]*entry{},
		rmx: &sync.RWMutex{},
	}
	for i := 0; i < 10; i++ {
		m.Out("lidong", i)
	}
	rd := m.Rd("lidong2", time.Second)
	rd = m.Rd("lidong", time.Second)
	fmt.Println(`rd := m.Rd("lidong", time.Second): `, rd)
	for {
	}
}
