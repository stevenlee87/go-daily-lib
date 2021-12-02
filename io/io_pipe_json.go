package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type PayLoad struct {
	Content string
}

func main() {

	r, w := io.Pipe()

	go func() {
		defer w.Close()

		err := json.NewEncoder(w).Encode(&PayLoad{Content: "Hello there!"})

		if err != nil {
			log.Fatal(err)
		}
	}()

	resp, err := http.Post("https://httpbin.org/post", "application/json", r)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}

/*
The example posts a JSON payload to the website and reads its body response.

go func() {
    defer w.Close()

    err := json.NewEncoder(w).Encode(&PayLoad{Content: "Hello there!"})

    if err != nil {
        log.Fatal(err)
    }
}()
We write a JSON payload to the PipeWriter in a goroutine.

resp, err := http.Post("https://httpbin.org/post", "application/json", r)
The http.Post method expects a reader as its third parameter; we pass a PipeReader there.

body, err := ioutil.ReadAll(resp.Body)

if err != nil {
    log.Fatal(err)
}

fmt.Println(string(body))
Finally, we read the response body and print it.

{
  "args": {},
  "data": "{\"Content\":\"Hello there!\"}\n",
  "files": {},
  "form": {},
  "headers": {
    "Accept-Encoding": "gzip",
    "Content-Type": "application/json",
    "Host": "httpbin.org",
    "Transfer-Encoding": "chunked",
    "User-Agent": "Go-http-client/2.0",
    "X-Amzn-Trace-Id": "Root=1-61a86e84-62a7ce7c7374e48e6f087da6"
  },
  "json": {
    "Content": "Hello there!"
  },
  "origin": "192.168.1.5",
  "url": "https://httpbin.org/post"
}

*/
