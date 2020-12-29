package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

/*
Unmarshal
既然有Marshal，也就有它的相反操作Unmarshal。Unmarshal用于将一个[]byte类型转换为相应的结构体或其他类型。
与Marshal不同，Unmarshal需要知道转换的目标类型，所以需要提供一个函数用于生成该类型的对象。然后将[]byte数据Unmarshal到该对象中。
Unmarshal接受两个参数，参数一是类型为func([]byte, interface{}) error的函数，参数二是func () interface{}用于生成实际类型的对象。
我们拿上面的例子中生成的 JSON 字符串作为数据，将它们重新Unmarshal为User对象：
*/

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	observable := rxgo.Just(
		`{"name":"syf","age":18}`,
		`{"name":"xy","age":20}`,
	)()

	observable = observable.Map(func(_ context.Context, i interface{}) (interface{}, error) {
		return []byte(i.(string)), nil
	}).Unmarshal(json.Unmarshal, func() interface{} {
		return &User{}
	})

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}

/*
由于Unmarshal接受[]byte类型的参数，我们在Unmarshal之前加了一个Map用于将string转为[]byte。运行：

$ go run main.go
&{syf 18}
&{xy 20}
*/
