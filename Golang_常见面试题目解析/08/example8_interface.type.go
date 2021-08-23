package main

type student struct {
	Name string
}

func jay(v interface{}) {
	switch msg := v.(type) {
	case *student, student:
		msg.Name
	}
}
func main() {
	jay("jay")
}

/*
golang中有规定， switch type 的 case T1 ，类型列表只有⼀个，那么 msg := v.(type)中的 v 的类型就是T1类型。
如果是 case T1, T2 ，类型列表中有多个，那 v 的类型还是多对应接⼝的类型，也就是 m 的类型。

所以这⾥ msg 的类型还是 interface{} ，所以他没有 Name 这个字段，编译阶段就会
报错。具体解释⻅： https://golang.org/ref /spec#Type_switches
*/
