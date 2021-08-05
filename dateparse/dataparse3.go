package main

import (
	"fmt"
	"time"

	"github.com/araddon/dateparse"
)

/*
dateparse支持在特定时区解析日期时间字符串。我们可以通过调用标准库的time.LoadLocation()方法，传入时区标识字符串来获得时区对象。
时区标识字符串是类似Asia/Shanghai，America/Chicago这样的格式，它表示一个具体的时区，前者上海，后者洛杉矶。
调用dateparse.ParseIn()方法传入时区对象，在指定时区中解析。
time包中还预定义了两个时区对象，time.Local表示本地时区，time.UTC表示 UTC 时区。时区的权威数据请看IANA。
*/

func main() {
	tz1, _ := time.LoadLocation("America/Chicago")
	t1, _ := dateparse.ParseIn("2021-08-05 14:56:30", tz1)
	fmt.Println(t1.Local().Format("2006-01-02 15:04:05"))

	t2, _ := dateparse.ParseIn("2021-08-05 14:56:30", time.Local)
	fmt.Println(t2.Local().Format("2006-01-02 15:04:05"))
}
