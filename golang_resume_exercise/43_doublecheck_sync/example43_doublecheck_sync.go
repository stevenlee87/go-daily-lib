package main

import (
	"fmt"
	"sync"
)

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if o.done == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.done = 1
		f()
	}
}

func main() {
	// 结构体初始化
	o := Once{m: sync.Mutex{},
		done: 0}
	o.Do(func() { fmt.Println("ok") })
}

/*
结构体初始化
type Map struct {
	c   map[string]*entry
	rmx *sync.RWMutex
}

	m := Map{
		c:   map[string]*entry{},
		rmx: &sync.RWMutex{},
	}

*/
