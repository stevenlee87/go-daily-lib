package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"time"
)

func main() {
	c := cron.New(
		cron.WithLogger(
			cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))
	c.AddFunc("@every 1s", func() {
		fmt.Println("hello world")
	})
	c.Start()

	time.Sleep(5 * time.Second)
}
