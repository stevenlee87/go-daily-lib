package main

type Student struct {
	name string
}

func main() {
	m := map[string]Student{"people": {"zhoujielun"}}
	m["people"].name = "wuyanzu"
}

// cannot assign to struct field m["people"].name in map
/*
解析：
map的value本身是不可寻址的，因为map中的值会在内存中移动，并且旧的指针地址在
map改变时会变得⽆效。故如果需要修改map值，可以将 map 中的⾮指针类型
value ，修改为指针类型，⽐如使⽤ map[string]*Student .
*/
