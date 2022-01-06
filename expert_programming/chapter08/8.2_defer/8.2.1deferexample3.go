package main

import "fmt"

func main() {
	var aArray = [3]int{1, 2, 3}
	defer func(array *[3]int) {
		for i := range array {
			fmt.Print(array[i])
		}
	}(&aArray)
	aArray[0] = 10
}

/*
OutPut:
10, 2, 3
延迟函数通过参数绑定了数组的地址，defer语句执行时访问的是更改之后的数组。
*/
