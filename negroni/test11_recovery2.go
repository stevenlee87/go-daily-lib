package main

import (
	"net/http"

	"github.com/urfave/negroni"
)

/*
这在开发环境比较有用，但是生产环境中不能泄露这个信息。这时可以设置PrintStack字段为false：
*/
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		panic("internal server error")
	})

	n := negroni.New()
	r := negroni.NewRecovery()
	r.PrintStack = false
	n.Use(r)
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
