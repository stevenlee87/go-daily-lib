package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

/*
mapstructure用于将通用的map[string]interface{}解码到对应的 Go 结构体中，或者执行相反的操作。很多时候，解析来自多种源头的数据流时，
我们一般事先并不知道他们对应的具体类型。只有读取到一些字段之后才能做出判断。这时，我们可以先使用标准的encoding/json库将数据解码
为map[string]interface{}类型，然后根据标识字段利用mapstructure库转为相应的 Go 结构体以便使用。
*/
type Person struct {
	Name string
	Age  int
	Job  string
}

type Cat struct {
	Name  string
	Age   int
	Breed string
}

func main() {
	datas := []string{`
    { 
      "type": "person",
      "name":"steven",
      "age":18,
      "job": "programmer"
    }
  `,
		`
    {
      "type": "cat",
      "name": "pongki",
      "age": 5,
      "breed": "Ragdoll"
    }
  `,
	}

	for _, data := range datas {
		var m map[string]interface{}
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}

		switch m["type"].(string) {
		case "person":
			var p Person
			mapstructure.Decode(m, &p)
			fmt.Println("person", p)

		case "cat":
			var cat Cat
			mapstructure.Decode(m, &cat)
			fmt.Println("cat", cat)
		}
	}
}
