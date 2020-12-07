package main

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

func main() {
	jObj, _ := gabs.ParseJSON([]byte(`{"info":{"age":18,"name":{"first":"steven","last":"lee"},"hobby":"programming"}}`))

	fmt.Println(jObj.String())
	jObj.Delete("info", "name")
	fmt.Println(jObj.String())

	jObj.Delete("info", "hobby")
	fmt.Println(jObj.String())
}
