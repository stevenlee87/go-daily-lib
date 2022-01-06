package main

import (
	"fmt"
	"os"
)

/*
在介绍基础error时，我们提到检查error时往往还会做类型转换，即使用断言来检测error是否是指定类型，比如：
if e, ok := err2.(*os.PathError); ok {
		fmt.Printf("it's an os.PathError, operation: %s, path: %s, msg: %v\n", e.Op, e.Path, e.Err)
	}

类型地，如果err是一个链式error，那么上面的代码将不再有效，比如下面的Example将没有任何打印结果：
*/
func ExampleAssetChanErrorWithoutAs() {
	err := &os.PathError{
		Op:   "write",
		Path: "/root/demo.txt",
		Err:  os.ErrPermission,
	}

	err2 := fmt.Errorf("some context: %w", err)
	if e, ok := err2.(*os.PathError); ok {
		fmt.Printf("it's an os.PathError, operation: %s, path: %s, msg: %v\n", e.Op, e.Path, e.Err)
	}
}

func main() {
	ExampleAssetChanErrorWithoutAs()
}
