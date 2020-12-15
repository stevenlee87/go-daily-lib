package main

import (
	"net/http"

	"github.com/urfave/negroni"
)

/*
设置PanicHandlerFunc之后，发生panic就会调用此函数。
我们还可以对输出的格式进行设置，设置Formatter字段为negroni.HTMLPanicFormatter能让输出更好地在浏览器中呈现：
*/

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		panic("internal server error")
	})

	n := negroni.New()
	r := negroni.NewRecovery()
	r.Formatter = &negroni.HTMLPanicFormatter{}
	n.Use(r)
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
