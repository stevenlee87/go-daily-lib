package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func NetWorkStatus() bool {
	cmd := exec.Command("ping", "www.baidu.com", "-c", "4", "-W", "5")
	cmd.Stdout = os.Stdout
	fmt.Println("NetWorkStatus Start:", time.Now())
	err := cmd.Run()
	fmt.Println("NetWorkStatus End  :", time.Now())

	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		fmt.Println("Net Status , OK")
	}
	return true
}

func main() {
	NetWorkStatus()
}
