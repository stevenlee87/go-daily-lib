package gotest

import "fmt"

// SayHello 打印一行字符串
func SayHello() {
	fmt.Println("Hello World")
}

// SayGoodbye 打印两行字符串
func SayGoodbye() {
	fmt.Println("Hello,")
	fmt.Println("goodbye")
}

// PrintNames 打印学生姓名
func PrintNames() {
	students := make(map[int]string, 4)
	students[1] = "Jim"
	students[2] = "Bob"
	students[3] = "Tom"
	students[4] = "Sue"
	for _, value := range students {
		fmt.Println(value)
	}
}

/*
这几个方法打印内容略有不同，分别代表一种典型的场景：
SayHello()：只有一行打印输出
SayGoodbye()：有两行打印输出
PrintNames()：有多行打印输出，且由于Map数据结构的原因，多行打印次序是随机的。
*/
