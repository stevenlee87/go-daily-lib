# Chapter 7 测试

目前go test支持的测试类型有：
- 单元测试
- 性能测试
- 示例测试

## 7.1 快速开始

1.单元测试  
单元测试是指对软件中的最小测试单元进行检测和验证，比如对一个函数的测试。

2.性能测试  
性能测试也成为基准测试，可以测试一段程序的性能，可以得到时间消耗、内存使用情况的报告。

3.示例测试，广泛应用于Go源码和各种开源框架中，用于展示某个包或某个方法的用法。 

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