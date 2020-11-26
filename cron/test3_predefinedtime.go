package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	c := cron.New()

	c.AddFunc("@hourly", func() {
		fmt.Println("Every hour")
	})

	c.AddFunc("@daily", func() {
		fmt.Println("Every day on midnight")
	})

	c.AddFunc("@weekly", func() {
		fmt.Println("Every week")
	})

	c.Start()

	for {
		time.Sleep(time.Second)
	}
}
