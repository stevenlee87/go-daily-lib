package __1_quick_start_test

import (
	"testing"
	"time"

	"github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/7.1_quick_start/gotest"
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

func BenchmarkSetBytes(b *testing.B) {
	b.SetBytes(1024 * 1024)
	for i := 0; i < b.N; i++ {
		time.Sleep(1 * time.Second) // 模拟待测函数
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("A=1", BenchmarkMakeSliceWithoutAlloc)
	b.Run("A=2", BenchmarkMakeSliceWithPreAlloc)
	b.Run("B=1", BenchmarkSetBytes)
}
