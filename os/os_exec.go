package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	cmd := exec.Command("ping", "baidu.com")
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}

	cmd.Start()

	buf := bufio.NewReader(stdout)
	num := 0

	for {
		line, _, _ := buf.ReadLine()
		if num > 3 {
			os.Exit(0)
		}
		num += 1
		fmt.Println(string(line))
	}
}

/*
Go cmd StdoutPipe
The StdoutPipe of a command returns a pipe that will be connected to the command's standard output when the command starts.

In the example, we launch a ping command and read four lines of its output.

cmd := exec.Command("ping", "webcode.me")
We create a command which lauches ping to test the availability of the webcode.me website.

stdout, err := cmd.StdoutPipe()
We get the standard output of the command.

buf := bufio.NewReader(stdout)
A reader from the standard output is created.

for {
    line, _, _ := buf.ReadLine()
    if num > 3 {
        os.Exit(0)
    }
    num += 1
    fmt.Println(string(line))
}
In a for loop, we read four lines and print them to the console.
PING webcode.me (46.101.248.126): 56 data bytes
64 bytes from 46.101.248.126: icmp_seq=0 ttl=43 time=308.692 ms
64 bytes from 46.101.248.126: icmp_seq=1 ttl=43 time=308.064 ms
64 bytes from 46.101.248.126: icmp_seq=2 ttl=43 time=308.081 ms
This is a sample output.

*/
