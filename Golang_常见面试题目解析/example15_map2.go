package main

import "fmt"

type Student struct {
	name string
}

func main() {
	m := map[string]*Student{"people": {"zhoujielun"}}
	m["people"].name = "wuyanzu"
	fmt.Println("m[\"people\"].name is:", m["people"].name)
}

// cannot assign to struct field m["people"].name in map
