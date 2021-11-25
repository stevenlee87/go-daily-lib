# Chapter 7 测试

目前go test支持的测试类型有：
- 单元测试
- 性能测试
- 示例测试

## 7.1 快速开始

### 1.单元测试  
单元测试是指对软件中的最小测试单元进行检测和验证，比如对一个函数的测试。

### 2.性能测试  
性能测试也成为基准测试，可以测试一段程序的性能，可以得到时间消耗、内存使用情况的报告。

### 3.示例测试  
广泛应用于Go源码和各种开源框架中，用于展示某个包或某个方法的用法。  
比如，Go标准库中，mail包展示如何从一个字符串解析出邮件列表的用法，非常直观易懂。  
源码位于 src/net/mail/example_test.go 中：  
```go
func ExampleParseAddressList() {
    const list = "Alice <alice@example.com>, Bob <bob@example.com>, Eve <eve@example.com>"
    emails, err := mail.ParseAddressList(list)
    if err != nil {
        log.Fatal(err)
    }

    for _, v := range emails {
        fmt.Println(v.Name, v.Address)
    }

    // Output:
    // Alice alice@example.com
    // Bob bob@example.com
    // Eve eve@example.com
}
```

## 总结
### 1. 单元测试  
通过package语句可以看到，测试文件属于“gotest_test”包，测试文件也可以跟源文件在同一个包，但常见的做法是创建一个包专用于测试，
这样可以使测试文件和源文件隔离。GO源代码以及其他知名的开源框架通常会创建测试包，而且规则是在原包名上加上”_test”。

测试函数命名规则为”TestXxx”，其中“Test”为单元测试的固定开头，go test只会执行以此为开头的方法。
紧跟“Test”是以首字母大写的单词，用于识别待测试函数。

测试函数参数并不是必须要使用的，但”testing.T”提供了丰富的方法帮助控制测试流程。

t.Errorf()用于标记测试失败，标记失败还有几个方法，在介绍testing.T结构时再详细介绍。 

### 2. 性能测试
性能测试函数命名规则为”BenchmarkXxx”，其中”Xxx”为自定义的标识，需要以大写字母开始，通常为待测函数。

testing.B提供了一系列的用于辅助性能测试的方法或成员，比如本例中的 b.N 表示循环执行的次数，而N值不用程序员特别关心，按照官方说法，
N值是动态调整的，直到可靠的算出程序执行时间后才会停止，具体执行次数会在执行结束后打印出来。

其中 -bench 为go test的flag，该flag指示go test进行性能测试，即执行测试文件中符合”BenchmarkXxx”规 则的方法。
紧跟flag的为flag的参数，本例表示执行当前所有的性能测试。

### 3. 示例测试
1. 例子测试函数名需要以”Example”开头；
2. 检测单行输出格式为“// Output: <期望字符串>”；
3. 检测多行输出格式为“// Output: \ <期望字符串> \ <期望字符串>”，每个期望字符串占一行；
4. 检测无序输出格式为”// Unordered output: \ <期望字符串> \ <期望字符串>”，每个期望字符串占一 行；
5. 测试字符串时会自动忽略字符串前后的空白字符；
6. 如果测试函数中没有“Output”标识，则该测试函数不会被执行；
7. 执行测试可以使用 go test ，此时该目录下的其他测试文件也会一并执行；
8. 执行测试可以使用 go test <xxx_test.go> ，此时仅执行特定文件中的测试函数；