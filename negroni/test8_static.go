package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

/*
1.在程序运行目录下创建public目录，然后放入一些文件1.txt，2.jpg。程序运行之后，就能通过浏览器localhost:3000/1.txt和
localhost:3000/2.jpg请求这些文件了。
2.另外需要特别注意一点，如果找不到对应的文件，Static会将请求传给下一个中间件或处理器函数。
3.在上面的例子中就是hello world。在浏览器中输入localhost:3000/none-exist.txt看看效果。
*/
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	n := negroni.New()
	n.Use(negroni.NewStatic(http.Dir("./public")))
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
