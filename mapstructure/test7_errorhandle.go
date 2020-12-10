package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

/*
错误处理
mapstructure执行转换的过程中不可避免地会产生错误，例如 JSON 中某个键的类型与对应 Go 结构体中的字段类型不一致。
Decode/DecodeMetadata会返回这些错误：
*/

type Person struct {
	Name   string
	Age    int
	Emails []string
}

func main() {
	m := map[string]interface{}{
		"name":   123,
		"age":    "bad value",
		"emails": []int{1, 2, 3},
	}

	var p Person
	err := mapstructure.Decode(m, &p)
	if err != nil {
		fmt.Println(err.Error())
	}
}

/*
* 'Age' expected type 'int', got unconvertible type 'string', value: 'bad value'
* 'Emails[0]' expected type 'string', got unconvertible type 'int', value: '1'
* 'Emails[1]' expected type 'string', got unconvertible type 'int', value: '2'
* 'Emails[2]' expected type 'string', got unconvertible type 'int', value: '3'
* 'Name' expected type 'string', got unconvertible type 'int', value: '123'
 */
