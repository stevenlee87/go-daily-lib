package main

import "fmt"

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		fmt.Printf("i is: %d\n", i)

		left1, right1 := expandAroundCenter(s, i, i)
		fmt.Printf("left1 and right1 is: %d %d\n", left1, right1)

		left2, right2 := expandAroundCenter(s, i, i+1)
		fmt.Printf("left2 and right2 is: %d %d\n", left2, right2)

		if right1-left1 > end-start {
			start, end = left1, right1
			fmt.Printf("right1-left1 > end-start---> start and end is: %d %d\n", start, end)
		}

		if right2-left2 > end-start {
			start, end = left2, right2
			fmt.Printf("right2-left2 > end-start---> start and end is: %d %d\n", start, end)
		}

	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
		fmt.Printf("left and right is: %d %d\n", left, right)
	}
	fmt.Printf("last result: left and right is: %d %d\n", left, right)
	return left + 1, right - 1
}

func main() {
	//s := "babad"
	s := "aa"
	result := longestPalindrome(s)
	fmt.Printf("输入：s = %s  输出：%s\n", s, result)
}
