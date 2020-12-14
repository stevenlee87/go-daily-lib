package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

/*
negroni的使用非常简单，它可以很方便的与http.Handler一起使用。negroni.Classic()提供了几个常用的中间件：

negroni.Recovery：恢复panic，处理器代码中有panic会被这个中间件捕获，程序不会退出；
negroni.Logger：日志，记录请求和响应的基本信息；
negroni.Static：在public目录提供静态文件服务。
调用n.UseHandler(mux)，将这些中间件应用到多路复用器上。运行，在浏览器中输入localhost:3000，查看控制台输出：
*/
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
