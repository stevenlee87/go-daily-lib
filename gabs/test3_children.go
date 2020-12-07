package main

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

/*
gabs提供了两个方法可以方便地遍历数组和对象：

Children()：返回所有数组元素的切片，如果在对象上调用该方法，Children()将以不确定顺序返回对象所有的值的切片；
ChildrenMap()：返回对象的键和值。
*/

func main() {
	jObj, _ := gabs.ParseJSON([]byte(`{
    "user": {
      "name": "steven",
      "age": 18,
      "members": [
        {
          "name": "hxy",
          "age": 20,
          "relation": "son"
        },
        {
          "name": "syf",
          "age": 20,
          "relation": "brother"
        }
      ]
    }
  }`))

	for k, v := range jObj.S("user").ChildrenMap() {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}

	fmt.Println()

	for i, v := range jObj.S("user", "members", "*").Children() {
		fmt.Printf("member %d: %v\n", i+1, v)
	}
}
