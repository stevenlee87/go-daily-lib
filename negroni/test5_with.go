package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

/*
negroni.With
如果有多个中间件，每个都需要n.Use()有些繁琐。negroni提供了一个With()方法，它接受一个或多个negroni.Handler参数，返回一个新的对象：
*/

func Middleware1(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("Middleware1 begin")
	next(w, r)
	fmt.Println("Middleware1 end")
}

func Middleware2(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("Middleware2 begin")
	next(w, r)
	fmt.Println("Middleware2 end")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	n := negroni.New()
	n = n.With(
		negroni.HandlerFunc(Middleware1),
		negroni.HandlerFunc(Middleware2),
	)
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
