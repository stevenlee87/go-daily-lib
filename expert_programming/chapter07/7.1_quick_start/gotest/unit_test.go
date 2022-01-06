package gotest_test

import (
	"testing"

	"github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start/gotest"
)

func TestAdd(t *testing.T) {
	var a = 1
	var b = 2
	var expected = 3

	actual := gotest.Add(a, b)
	if actual != expected {
		t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expected)
	}
}

/* ~/project/golang/go-daily-lib/expert_programming/chapter7/gotest$ go test
PASS
ok      github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/gotest  1.224s

通过package语句可以看到，测试文件属于“gotest_test”包，测试文件也可以跟源文件在同一个包，但常见的做 法是创建一个包专用于测试，
这样可以使测试文件和源文件隔离。GO源代码以及其他知名的开源框架通常会创建测试包，而且规则是在原包名上加上”_test”。

测试函数命名规则为”TestXxx”，其中“Test”为单元测试的固定开头，go test只会执行以此为开头的方法。
紧跟“Test”是以首字母大写的单词，用于识别待测试函数。

测试函数参数并不是必须要使用的，但”testing.T”提供了丰富的方法帮助控制测试流程。

t.Errorf()用于标记测试失败，标记失败还有几个方法，在介绍testing.T结构时再详细介绍。
*/
