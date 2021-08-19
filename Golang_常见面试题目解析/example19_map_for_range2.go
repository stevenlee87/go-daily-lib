package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		value := stu
		m[stu.Name] = &value
		fmt.Println("m[stu.Name] is:", m[stu.Name])
	}

	fmt.Println(m["zhou"])
	fmt.Println(m["li"])
	fmt.Println(m["wang"])
}
