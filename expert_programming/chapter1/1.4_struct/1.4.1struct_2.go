package main

import "fmt"

type Kid struct {
	Name string
	Age  int
}

func (k Kid) SetName(name string) {
	k.Name = name
}

func (k *Kid) SetAge(age int) {
	k.Age = age
}

func (k *Kid) GetName() string {
	return k.Name
}

func main() {
	k := Kid{Name: "lee", Age: 18}
	k.SetName("huang")
	fmt.Println(k.GetName())
}

/*
SetName 无法修改名字
*/
