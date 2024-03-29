package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

/*
除了在控制台和浏览器中输出panic信息，Recovery还提供了钩子函数，可以向其他服务上报panic，如Sentry/Airbrake。当然上报的代码要自己写
*/
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		panic("internal server error")
	})

	n := negroni.New()
	r := negroni.NewRecovery()
	r.PanicHandlerFunc = reportToSentry
	n.Use(r)
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}

func reportToSentry(info *negroni.PanicInformation) {
	fmt.Println("sent to sentry")
}
