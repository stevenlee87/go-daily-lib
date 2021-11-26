package gotest_test

import gotest "github.com/stevenlee87/go-daily-lib/expert_programming/chapter7/example"

// 检测单行输出
func ExampleSayHello() {
	gotest.SayHello()
	// OutPut: Hello World
}

// 检测多行输出
func ExampleSayGoodbye() {
	gotest.SayGoodbye()
	// OutPut:
	// Hello,
	// goodbye
}

// 检测乱序输出
func ExamplePrintNames() {
	gotest.PrintNames()
	// Unordered output:
	// Jim
	// Bob
	// Tom
	// Sue
}

/*
例子测试函数命名规则为”Examplexxx”，其中”xxx”为自定义的标识，通常为待测函数名称。
这三个测试函数分别代表三种场景：
ExampleSayHello()： 待测试函数只有一行输出，使用”// OutPut: “检测。
ExampleSayGoodbye()：待测试函数有多行输出，使用”// OutPut: “检测，其中期望值也是多行。
ExamplePrintNames()：待测试函数有多行输出，但输出次序不确定，使用”// Unordered output:”检测。
注：字符串比较时会忽略前后的空白字符。

~/project/golang/go-daily-lib/expert_programming/chapter7/example$ go test example_test.go
ok      command-line-arguments  0.515s

1. 例子测试函数名需要以”Example”开头；
2. 检测单行输出格式为“// Output: <期望字符串>”；
3. 检测多行输出格式为“// Output: \ <期望字符串> \ <期望字符串>”，每个期望字符串占一行；
4. 检测无序输出格式为”// Unordered output: \ <期望字符串> \ <期望字符串>”，每个期望字符串占一 行；
5. 测试字符串时会自动忽略字符串前后的空白字符；
6. 如果测试函数中没有“Output”标识，则该测试函数不会被执行；
7. 执行测试可以使用 go test ，此时该目录下的其他测试文件也会一并执行；
8. 执行测试可以使用 go test <xxx_test.go> ，此时仅执行特定文件中的测试函数；
*/
