package main

import (
	"sync"
	"time"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	ua.Lock()
	defer ua.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	ua := UserAges{
		ages:  map[string]int{"liDong": 18},
		Mutex: sync.Mutex{},
	}

	go func() {
		for i := 0; i < 100; i++ {
			ua.Add(time.Now().String(), time.Now().Second())

		}
	}()
	for i := 0; i < 100; i++ {
		ua.Get(time.Now().String())
	}
}

// go run -race example25_struct2.go
