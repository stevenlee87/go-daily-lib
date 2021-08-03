package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "Hello" + " " + "World"
	fmt.Println(s1)

	ss := []string{"Hello", "World"}
	fmt.Println(strings.Join(ss, " "))
}

/*
需要注意的是，将待拼接的字符串放在一行中，使用+拼接，在 Go 语言内部会先计算需要的空间，预先分配这个空间，最后将各个字符串拷贝过去。
这个行为与其他很多语言是不同的，所以在 Go 语言中使用+拼接字符串不会有性能损失，甚至由于内部优化比其他方式性能还要更好一些。
当然前提拼接是一次完成的。
*/
