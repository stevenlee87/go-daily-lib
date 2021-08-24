package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	Name string `json:"name"` //TODO name
	Sex  string `json:"sex"`
}

type Student struct {
	Name string `json:"name"` //TODO name
	Id   string `json:"id"`
}

type User struct {
	*Person
	*Student // TODO name会重复.
}

func main() {
	str := `{
  "name": "张三",
  "sex" : "男",
  "Id": "10001"
}`

	user := new(User)
	err := json.Unmarshal([]byte(str), &user)
	log.Println("ERR:", err)
	str1, err := json.Marshal(user)
	log.Println("user json==>", string(str1))
}

/*
2021/08/24 14:35:44 ERR: <nil>
2021/08/24 14:35:44 user json==> {"sex":"男","id":"10001"}

User 内嵌两个匿名结构体Person 和Student ，Person 和 Student 具有相同的字段Name，
那么User的实例（user）是不能直接访问Name字段。
你定义的类型属于成员名字冲突，必须通过user.Person.Name 或者user.Student.Name
才能具体确定属于谁

结构内嵌特性:
1、内嵌的结构体可以直接访问其成员变量
2、内嵌结构体的字段名是它的类型名
要实现你的需求，可以在User添加Name字段
*/
