package main

import (
	"fmt"
	cregex "github.com/mingrammer/commonregex"
)

func main() {
	text := `John, please get that article on www.linkedin.com to me by 5:00PM on Jan 9th 2012. 4:00 would be ideal, 
actually. If you have any questions, You can reach me at 15210012345 or get in touch with 
my associate at harold.smith@gmail.com`

	dateList := cregex.Date(text)
	timeList := cregex.Time(text)
	linkList := cregex.Links(text)
	phoneList := cregex.Phones(text)
	emailList := cregex.Emails(text)

	fmt.Println("date list:", dateList)
	fmt.Println("time list:", timeList)
	fmt.Println("link list:", linkList)
	fmt.Println("phone list:", phoneList)
	fmt.Println("email list:", emailList)
}

/*
date list: [Jan 9th 2012]
time list: [5:00PM 4:00 ]
link list: [www.linkedin.com harold.smith@gmail.com]
phone list: [15210012345]
email list: [harold.smith@gmail.com]
*/
