package main

import (
	"fmt"
	"log"
)

func fire(f func(i ...int), i ...int) {
	f(i...)
	log.Printf("%T", f)
}

func f1(i ...int) {
	log.Println("f1 p: ", i)
}

func f2(i ...int) {
	log.Println("f2 p: ", i)
}

func f3(i ...int) {
	log.Println("f3 p: ", i)
}

func main() {
	var (
		fm = map[string]func(i ...int){
			"f1": f1,
			"f2": f2,
			"f3": f3,
		}
	)

	for _, v := range fm {
		fire(v, 10, 11, 12, 13, 14, 15)
		fmt.Println()
	}
}
