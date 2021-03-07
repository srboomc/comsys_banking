package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// log.Println("chen")
		netData, err := bufio.NewReader(c).ReadString('\n')
		// log.Println("chen1")
		if err != nil {
			// log.Println("chen2")
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			// log.Println("chen3")
			fmt.Println("Exiting TCP server!")
			return
		}
		// log.Println("chen4")
		fmt.Print("-> ", string(netData))
		// log.Println("chen5")
		t := time.Now()
		// log.Println("chen6")
		myTime := t.Format(time.RFC3339) + "\n"
		// log.Println("chen7")
		c.Write([]byte(myTime))
		// log.Println("chen8")
	}
}
