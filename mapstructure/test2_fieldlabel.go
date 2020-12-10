package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string `mapstructure:"username"`
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
      "username":"steven",
      "age":18,
      "job": "programmer"
    }
  `,
		`
    {
      "type": "cat",
      "name": "kitty",
      "Age": 1,
      "breed": "Ragdoll"
    }
  `,
		`
    {
      "type": "cat",
      "Name": "rooooose",
      "age": 2,
      "breed": "shorthair"
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

/*
上面代码中，我们使用标签mapstructure:"username"将Person的Name字段映射为username，在 JSON 串中我们需要设置username才能正确解析。
另外，注意到，我们将第二个 JSON 串中的Age和第三个 JSON 串中的Name首字母大写了，但是并没有影响解码结果。mapstructure处理字段映射是大小写
不敏感的。
*/
