package main

import (
	"fmt"
)

type Student struct {
	Name string
}

func main() {
	fmt.Println(&Student{Name: "jay"} == &Student{Name: "jay"})
	fmt.Println(Student{Name: "jay"} == Student{Name: "jay"})
	fmt.Println(Student{Name: "jay"})
}

// 指针类型⽐较的是指针地址，⾮指针类型⽐较的是每个属性的值。
