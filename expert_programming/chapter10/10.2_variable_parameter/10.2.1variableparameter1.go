package main

import "fmt"

func Greeting(prefix string, who ...string) {
	if who == nil {
		fmt.Printf("Nobody to say hi.\n\n")
		return
	}

	for _, people := range who {
		fmt.Printf("%s %s\n", prefix, people)
	}
}

func ExampleGreetingWithoutParameter() {
	Greeting("nobody")
}

func ExampleGreetingWithParameter() {
	Greeting("hello:", "Joe", "Anna", "Eileen")
}

func ExampleGreetingWithoutSlice() {
	guest := []string{"Joe", "Anna", "Eileen"}
	Greeting("hello:", guest...)
}

func main() {
	ExampleGreetingWithoutParameter()
	ExampleGreetingWithParameter()
	ExampleGreetingWithoutSlice()
}

/*
这个函数几乎把可变参函数的特征全部表现出来了：
可变参数必须在函数参数列表的尾部，即最后一个（如放前面会引起编译时歧义）；
可变参数在函数内部是作为切片来解析的；
可变参数可以不填，不填时函数内部当成 nil 切片处理；
可变参数必须是相同类型的（如果需要是不同类型的可以定义为interface{}类型）；
*/
