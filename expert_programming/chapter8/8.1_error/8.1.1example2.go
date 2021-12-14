package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	err := errors.New("not found")
	err1 := fmt.Errorf("some context: %v", err)
	err2 := fmt.Errorf("some context: %w", err)

	//fmt.Println(err1)
	//fmt.Println(err2)
	//some context: not found
	//some context: not found

	if err1 == err2 {
		fmt.Printf("two errors are equal\n")
	} else {
		fmt.Printf("two errors are different\n")
	}

	if err1.Error() != err2.Error() {
		panic("two errors's content should be same")
	}

	fmt.Printf("%s\n", reflect.TypeOf(err1).String())
	fmt.Printf("%s\n", reflect.TypeOf(err2).String())

	fmt.Println("err1.Error() is:", err1.Error())
	fmt.Println("err2.Error() is:", err2.Error())
}

/*
fmt.Errorf()的格式动词%v最终会调用err的Error()方法获取错误字符串，然后和some content:合并，生成errors包的errorString实例，而格动词%w
会生成fmt包的wrapError实例，所以二者的类型不同。
*/
