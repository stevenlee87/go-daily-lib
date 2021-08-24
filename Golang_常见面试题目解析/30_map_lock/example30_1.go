package main

import (
	"sync"
)

/*
实现阻塞读且并发安全的map GO⾥⾯MAP如何实现key不存在 get操作等待 直到key存在或者超时，保证并发安全，
且需要实现以下接⼝：
type sp interface {
	Out(key string, val interface{}) //存入key /val，如果该key读取的 goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回

	Rd(key string, timeout time.Duration) interface{} //读取一个key，如果 key不存在阻塞，等待key存在或者超时
}
*/

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
	item, ok := m.c[key]
	if !ok {
		m.c[key] = &entry{
			value:   val,
			isExist: true,
		}
		return
	}
	item.value = val
	if !item.isExist {
		if item.ch != nil {
			close(item.ch)
			item.ch = nil
		}
	}
	return
}
