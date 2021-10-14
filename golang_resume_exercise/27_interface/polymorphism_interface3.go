package main

import "fmt"

type Animal struct {
	Name string
}

type Cat struct {
	Animal
	FeatureA string
}

type Dog struct {
	Animal
	FeatureB string
}

type AnimalSounder interface {
	MakeDNA()
}

func MakeSomeDNA(animalSounder AnimalSounder) {
	animalSounder.MakeDNA()
}

func (c *Cat) MakeDNA() {
	fmt.Println("you are you!")
}

func (c *Dog) MakeDNA() {
	fmt.Println("You are not you!")
}

func main() {
	MakeSomeDNA(&Cat{})
	MakeSomeDNA(&Dog{})
}

/*
当 Cat 和 Dog 的实例实现了 AnimalSounder 接口类型的约束后，就意味着满足了条件，
他们在 Go 语言中就是一个东西。能够作为入参传入 MakeSomeDNA 方法中，
再根据不同的实例实现多态行为。

1.接口由使用者定义，使用者是个小孩，duck 就是个大黄鸭。
2.接口的实现是隐式的，不需要申明我实现了这个Retriever接口，只要实现接口里的方法
go语言就是一种duck typing
interface 需要完成的目的是实现了什么功能。至于具体怎么实现的使用者并不关心。接口是隐式的，没有继承关系。
Go语言实现一个接口并不需要显示声明，而是只要你实现了接口中的所有方法就认为你实现了这个接口。这称之为Duck typing。
*/
