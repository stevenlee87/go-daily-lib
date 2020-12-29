package main

import (
	"encoding/json"
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

/*
Marshal对经过它的数据进行一次Marshal。这个Marshal可以是json.Marshal/proto.Marshal，甚至我们自己写的Marshal函数。
它接受一个类型为func(interface{}) ([]byte, error)的函数用于对数据进行处理。
*/
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	observable := rxgo.Just(
		User{
			Name: "syf",
			Age:  18,
		},
		User{
			Name: "hxy",
			Age:  20,
		},
	)()

	observable = observable.Marshal(json.Marshal)

	for item := range observable.Observe() {
		fmt.Println(string(item.V.([]byte)))
	}
}

/*
{"name":"syf","age":18}
{"name":"hxy","age":20}
*/
