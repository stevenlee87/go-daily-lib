package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

/*
前面我们都是将map[string]interface{}解码到 Go 结构体中。mapstructure当然也可以将 Go 结构体反向解码为map[string]interface{}。在反
向解码时，我们可以为某些字段设置mapstructure:",omitempty"。这样当这些字段为默认值时，就不会出现在结构的map[string]interface{}中
*/
type Person struct {
	Name string
	Age  int
	Job  string `mapstructure:",omitempty"`
}

func main() {
	p := &Person{
		Name: "steven",
		Age:  18,
	}

	var m map[string]interface{}
	mapstructure.Decode(p, &m)

	data, _ := json.Marshal(m)
	fmt.Println(string(data))
}
