package gotest_test

import (
	"testing"
	"time"
)

// 并发子测试，无实际测试工作，仅用于演示
func parallelTest1(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
	// do some testing
}

// 并发子测试，无实际测试工作，仅用于演示
func parallelTest2(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
	// do some testing
}

//并发子测试，无实际测试工作，仅用于演示
func parallelTest3(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
	// do some testing
}

// TestSubParallel 通过把多个子测试放到一个组中并发执行，同时多个子测试可以共享setup和tear-down
func TestSubParallel(t *testing.T) {
	// setup
	t.Logf("Setup")

	t.Run("group", func(t *testing.T) {
		t.Run("Test1", parallelTest1)
		t.Run("Test2", parallelTest2)
		t.Run("Test3", parallelTest3)
	})

	// tear down
	t.Logf("teardown")
}

/*
setup 和 teardown 是在每个 case 执行前后都需要执行的操作

上面三个子测试中分别sleep了3s、2s、1s用于观察并发执行顺序。通过 Run() 将多个子测试“封装”到一个组中，
可以保证所有子测试全部执行结束后再执行tear-down。 命令行下的输出如下：

go test subparallel_test.go -v -run SubParallel
=== RUN   TestSubParallel
    subparallel_test.go:32: Setup
=== RUN   TestSubParallel/group
=== RUN   TestSubParallel/group/Test1
=== PAUSE TestSubParallel/group/Test1
=== RUN   TestSubParallel/group/Test2
=== PAUSE TestSubParallel/group/Test2
=== RUN   TestSubParallel/group/Test3
=== PAUSE TestSubParallel/group/Test3
=== CONT  TestSubParallel/group/Test1
=== CONT  TestSubParallel/group/Test2
=== CONT  TestSubParallel/group/Test3
=== CONT  TestSubParallel
    subparallel_test.go:41: teardown
--- PASS: TestSubParallel (3.00s)
    --- PASS: TestSubParallel/group (0.00s)
        --- PASS: TestSubParallel/group/Test3 (1.01s)
        --- PASS: TestSubParallel/group/Test2 (2.01s)
        --- PASS: TestSubParallel/group/Test1 (3.00s)
PASS
ok      command-line-arguments  4.280s

通过该输出可以看出：
1. 子测试是并发执行的（Test1最先被执行却最后结束）
2. tear-down在所有子测试结束后才执行
*/
