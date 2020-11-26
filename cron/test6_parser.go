package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

/*
const (
	Second         ParseOption = 1 << iota
	SecondOptional
	Minute
	Hour
	Dom
	Month
	Dow
	DowOptional
	Descriptor
)

var places = []ParseOption{
	Second,
	Minute,
	Hour,
	Dom,
	Month,
	Dow,
}

var defaults = []string{
	"0",
	"0",
	"0",
	"*",
	"*",
	"*",
}
*/

func main() {
	parser := cron.NewParser(
		cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)

	c := cron.New(cron.WithParser(parser))
	c.AddFunc("1 * * * * *", func() {
		fmt.Println("every 1 second")
	})
	c.Start()

	time.Sleep(62 * time.Second)
}
