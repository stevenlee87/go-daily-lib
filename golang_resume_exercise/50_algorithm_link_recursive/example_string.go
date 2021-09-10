package main

import (
	"fmt"
)

type Car struct {
	age  int
	name string
}

func (c Car) String() string {
	return fmt.Sprintf("Car-[name:%v, age:%v]", c.name, c.age)
}

func main() {
	c1 := Car{age: 1, name: "buick"}
	c2 := Car{age: 2, name: "benz"}
	fmt.Println(c1)
	// 对于切片中的对象，也能递归调用String()
	carSlice := []Car{}
	carSlice = append(carSlice, c1)
	carSlice = append(carSlice, c2)
	fmt.Println(carSlice)
}
