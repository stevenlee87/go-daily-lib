package main

import (
	"fmt"
	"strings"

	//"strings"

	"github.com/mingrammer/commonregex"
)

func main() {
	text := `mac address: ac:de:48:00:11:22, ip: 192.168.3.20, md5: fdbf72fdabb67ea6ef7ff5155a44def4 192.168.1.100 192.168.1.100`

	macList := commonregex.MACAddresses(text)
	ipList := commonregex.IPs(text)
	md5List := commonregex.MD5Hexes(text)

	fmt.Println("mac list:", macList)
	fmt.Println("ip list:", ipList)
	fmt.Println("md5 list:", md5List)

	for _, v := range ipList {
		//fmt.Println(v)
		if strings.Count(string(ipList), v) > 1 {
			fmt.Println(v)
		}
	}

}
