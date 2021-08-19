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
		m[stu.Name] = &stu
		fmt.Println("m[stu.Name] is:", m[stu.Name])
	}

	fmt.Println(m["zhou"])
	fmt.Println(m["li"])
	fmt.Println(m["wang"])
}

/*
m[stu.Name] is: &{zhou 24}
m[stu.Name] is: &{li 23}
m[stu.Name] is: &{wang 22}
&{wang 22}
&{wang 22}
&{wang 22}


解析：
golang 的 for ... range 语法中， stu 变量会被复⽤，每次循环会将集合中的值复制
给这个变量，因此，会导致最后 m 中的 map 中储存的都是 stus 最后⼀个 student的值
*/
