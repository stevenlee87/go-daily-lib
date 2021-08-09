package main

import (
	"fmt"
	"unicode"
)

const (
	Left = iota
	Top
	Right
	Bottom
)

func main() {
	//println(move("R2(LF)", 0, 0, Top))
	x1, y1, z1 := move("RF", 0, 0, Top)
	fmt.Println(x1, y1, z1)
}

func move(cmd string, x0 int, y0 int, z0 int) (x, y, z int) {
	x, y, z = x0, y0, z0
	fmt.Println("z:", z)
	repeat := 0
	repeatCmd := ""
	for _, s := range cmd {
		switch {
		case unicode.IsNumber(s):
			fmt.Println(int(s))
			r := int(s) - '0'
			fmt.Println(r)
			repeat = repeat*10 + (int(s) - '0')
		case s == ')':
			for i := 0; i < repeat; i++ {
				x, y, z = move(repeatCmd, x, y, z)
			}
			repeat = 0
			repeatCmd = ""
		case repeat > 0 && s != '(' && s != ')':
			repeatCmd = repeatCmd + string(s)
		case s == 'L':
			z = (z + 1) % 4
		case s == 'R':
			z = (z - 1 + 4) % 4
			fmt.Println(z)
		case s == 'F':
			switch {
			case z == Left || z == Right:
				x = x - z + 1
			case z == Top || z == Bottom:
				y = y - z + 2
			}
		case s == 'B':
			switch {
			case z == Left || z == Right:
				x = x + z - 1
			case z == Top || z == Bottom:
				y = y + z - 2
			}
		}
	}
	return
}

/*
机器⼈坐标问题
问题描述
有⼀个机器⼈，给⼀串指令，L左转 R右转，F前进⼀步，B后退⼀步，问最后机器⼈的
坐标，最开始，机器⼈位于 0 0，⽅向为正Y。 可以输⼊重复指令n ： ⽐如 R2(LF) 这
个等于指令 RLFLF。 问最后机器⼈的坐标是多少？

解题思路
这⾥的⼀个难点是解析重复指令。主要指令解析成功，计算坐标就简单了。
*/
