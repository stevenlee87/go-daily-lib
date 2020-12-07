package main

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

/*
合并
我们可以将两个gabs.Container合并成一个。如果同一个路径下有相同的键：

如果两者都是对象类型，则对二者进行合并操作；
如果两者都是数组类型，则将后者中所有元素追加到前一个数组中；
其中一个为数组，合并之后另一个同名键的值将会作为元素添加到数组中。
*/

func main() {
	obj1, _ := gabs.ParseJSON([]byte(`{"user":{"name":"steven"}}`))
	obj2, _ := gabs.ParseJSON([]byte(`{"user":{"age":18}}`))
	obj1.Merge(obj2)
	fmt.Printf("obj1:\n")
	fmt.Println(obj1)

	fmt.Printf("obj2:\n")
	arr1, _ := gabs.ParseJSON([]byte(`{"user": {"hobbies": ["game"]}}`))
	arr2, _ := gabs.ParseJSON([]byte(`{"user": {"hobbies": ["programming"]}}`))
	arr1.Merge(arr2)
	fmt.Println(arr1)

	fmt.Printf("obj3:\n")
	obj3, _ := gabs.ParseJSON([]byte(`{"user":{"name":"steven", "hobbies": "game"}}`))
	arr3, _ := gabs.ParseJSON([]byte(`{"user": {"hobbies": ["programming"]}}`))
	obj3.Merge(arr3)
	fmt.Println(obj3)

	fmt.Printf("obj4:\n")
	obj4, _ := gabs.ParseJSON([]byte(`{"user":{"name":"steven", "hobbies": "game"}}`))
	arr4, _ := gabs.ParseJSON([]byte(`{"user": {"hobbies": ["programming"]}}`))
	arr4.Merge(obj4)
	fmt.Println(arr4)

	fmt.Printf("obj5:\n")
	obj5, _ := gabs.ParseJSON([]byte(`{"user":{"name":"steven", "hobbies": {"first": "game"}}}`))
	arr5, _ := gabs.ParseJSON([]byte(`{"user": {"hobbies": ["programming"]}}`))
	obj5.Merge(arr5)
	fmt.Println(obj5)

	obj6, _ := gabs.ParseJSON([]byte(`{"user":{"name":"steven"}}`))
	obj7, _ := gabs.ParseJSON([]byte(`{"user":{"name":"lee"}}`))
	obj6.Merge(obj7)
	fmt.Printf("obj6:\n")
	fmt.Println(obj6)
	// {"user":{"name":["steven","lee"]}}

}
