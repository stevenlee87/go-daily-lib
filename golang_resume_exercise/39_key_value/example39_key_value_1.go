package main

import (
	"fmt"
)

type Student struct {
	Age int
}

func main() {
	kv := map[string]*Student{"menglu": {Age: 21}}

	for _, value := range kv {
		value.Age = 22
	}
	fmt.Println(kv["menglu"].Age)

}
