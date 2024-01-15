package main

import "fmt"

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
		fmt.Printf("left and right is: %d %d\n", left, right)
	}
	fmt.Printf("last result: left and right is: %d %d\n", left, right)
	return left + 1, right - 1
}

func main() {
	s := "babad"
	left1, right1 := expandAroundCenter(s, 2, 2)
	fmt.Printf("left1 and right1 is: %d %d\n", left1, right1)
}

/*
循环的执行过程如下：

1.执行初始化语句（如果有）。
2.检查循环条件，如果条件为 false，则退出循环。
3.执行循环体。
4.执行后置语句（如果有）。
5.返回步骤 2，继续下一次循环。
*/
