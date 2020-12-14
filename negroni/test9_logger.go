package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	n := negroni.New()
	logger := negroni.NewLogger()
	logger.SetFormat("[{{.Status}} {{.Duration}}] - {{.Request.UserAgent}}")
	n.Use(logger)
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}

/*
上面代码中将日志格式设置为[{{.Status}} {{.Duration}}] - {{.Request.UserAgent}}，
即响应状态、耗时和UserAgent。
*/
