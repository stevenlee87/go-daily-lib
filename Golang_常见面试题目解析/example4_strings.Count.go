package main

import (
	"fmt"
	"strings"
)

func isRegroup(s1, s2 string) bool {
	sl1 := len([]rune(s1))
	sl2 := len([]rune(s2))

	if sl1 > 5000 || sl2 > 5000 || sl1 != sl2 {
		return false
	}

	for _, v := range s1 {

		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}

func main() {
	b := isRegroup("ab", "ba")
	fmt.Println(b)
}

/*
判断两个给定的字符串排序后是否⼀致
问题描述
给定两个字符串，请编写程序，确定其中⼀个字符串的字符重新排列后，能否变成另⼀
个字符串。 这⾥规定【⼤⼩写为不同字符】，且考虑字符串重点空格。给定⼀个string s1和⼀个string s2，请返回⼀个bool，
代表两串是否重新排列后可相同。 保证两串的⻓度都⼩于等于5000。

解题思路
⾸先要保证字符串⻓度⼩于5000。之后只需要⼀次循环遍历s1中的字符在s2是否都存
在即可。
*/
