package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.LoadOrStore("a", 1)
	m.Delete("a")
	// m.Len undefined (type sync.Map has no field or method Len)
	//fmt.Println(m.Len()) // 源码中的Len根本就不存在

	fmt.Println(m.Load("a"))
}
