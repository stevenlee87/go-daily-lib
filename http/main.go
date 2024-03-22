package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	u := flag.String("url", "", "url to request")
	p := flag.Duration("pause", 0, "pause between requests")
	t := flag.Duration("threshold", 0, "threshold for reporting")
	c := flag.Int("count", 10000, "number of requests to send")
	n := flag.Int("close", 1, "close connection after every that many requests")

	flag.Parse()

	if *u == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	client := http.Client{}

	for i := 0; i < *c; i++ {
		started := time.Now()

		request, err := http.NewRequest("GET", *u, nil)
		if err != nil {
			log.Fatalf("Error constructing request: %s", err)
		}

		if i%*n == 0 {
			request.Header.Set("Connection", "Close")
		}

		response, err := client.Do(request)
		if err != nil {
			log.Fatalf("Error performing request: %s", err)
		}

		_, err = ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalf("Error reading request body: %s", err)
		}

		response.Body.Close()

		elapsed := time.Since(started)
		if elapsed > *t {
			log.Printf("Request %d took %dms", i, int(elapsed.Seconds()*1000))
		}

		time.Sleep(*p)
	}
}
