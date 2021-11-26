package gotest_test

import (
	"os"
	"testing"

	"github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start/gotest"
)

// TestAdd加法测试
func TestAdd(t *testing.T) {
	var a = 1
	var b = 2
	var expected = 3

	actual := gotest.Add(a, b)
	if actual != expected {
		t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expected)
	}
}

func BenchmarkMakeSliceWithoutAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotest.MakeSliceWithoutAlloc()
	}
}

// 检测多行输出
func ExampleSayGoodbye() {
	gotest.SayGoodbye()
	// OutPut:
	// Hello,
	// goodbye
}

// TestMain 用于主动执行各种测试，可以测试前后做setup和tear-down操作
func TestMain(m *testing.M) {
	println("TestMain setup.")

	retCode := m.Run()
	// 执行测试，包括单元测试、性能测试和示例测试

	println("TestMain tear-down.")

	os.Exit(retCode)
}

/*
单元测试：
go test . -v

性能测试：
go test -v -bench=.

示例测试：
go test main_test.go

*/
