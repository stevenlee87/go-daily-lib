package gotest_test

import (
	"testing"

	gotest "github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/benchmark"
)

func BenchmarkMakeSliceWithoutAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotest.MakeSliceWithoutAlloc()
	}
}

func BenchmarkMakeSliceWithPreAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotest.MakeSliceWithPreAlloc()
	}
}

/*
性能测试函数命名规则为”BenchmarkXxx”，其中”Xxx”为自定义的标识，需要以大写字母开始，通常为待测函数。

testing.B提供了一系列的用于辅助性能测试的方法或成员，比如本例中的 b.N 表示循环执行的次数，而N值不用程序员特别关心，按照官方说法，
N值是动态调整的，直到可靠的算出程序执行时间后才会停止，具体执行次数会在执行结束后打印出来。

go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/benchmark
cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz
BenchmarkMakeSliceWithoutAlloc-8            2047            502837 ns/op
BenchmarkMakeSliceWithPreAlloc-8            8690            160252 ns/op
PASS
ok      github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/benchmark       3.705s

其中 -bench 为go test的flag，该flag指示go test进行性能测试，即执行测试文件中符合”BenchmarkXxx”规 则的方法。
紧跟flag的为flag的参数，本例表示执行当前所有的性能测试。

通过输出可以直观的看出， BenchmarkMakeSliceWithoutAlloc 执行了2047次，平均每次502837纳秒，
BenchmarkMakeSliceWithPreAlloc 执行了8690次，平均每次160252纳秒。
从测试结果上看，虽然构造切片很快，但通过给切片预分配内存，性能还可以进一步提升，符合预期。关于原理分析，请参考Slice相关章节。
*/
