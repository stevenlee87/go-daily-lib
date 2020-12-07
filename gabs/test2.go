package main

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

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
          "age": 21,
          "relation": "brother"
        }
      ]
    }
  }`))

	fmt.Println("member names: ", jObj.S("user", "members", "*", "name").Data())
	fmt.Println("member ages: ", jObj.S("user", "members", "*", "age").Data())
	fmt.Println("member relations: ", jObj.S("user", "members", "*", "relation").Data())
	fmt.Println("son's name: ", jObj.S("user", "members", "0", "name").Data().(string))
}
