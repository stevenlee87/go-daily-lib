package main

import (
	"fmt"
	"log"

	"github.com/araddon/dateparse"
)

func main() {
	t1, err := dateparse.ParseAny("8/31/2021")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t1.Format("2006-01-02 15:04:05"))

	t2, err := dateparse.ParseAny("mm/dd/yyyy")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t2.Format("2006-01-02 15:04:05"))
}

/*
2021-08-31 00:00:00
2021/08/04 10:50:03 Could not find format for "mm/dd/yyyy"

dateparse默认采用mm/dd/yyyy这种格式
*/
