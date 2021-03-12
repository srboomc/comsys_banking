//referrence: https://stackoverflow.com/questions/51254367/making-golang-tcp-server-concurrent

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// only needed below for sample processing

func handleConnection(conn net.Conn) {
	fmt.Println("Inside function")
	// run loop forever (or until ctrl-c)
	// defer conn.Close()
	// for {
	fmt.Println("Inside loop")
	// will listen for message to process ending in newline (\n)
	message, _ := bufio.NewReader(conn).ReadString('\n')
	// output message received
	fmt.Print("Message Received:", string(message))
	// sample process for string received
	newmessage := strings.ToUpper(message)
	// send new string back to client
	conn.Write([]byte(newmessage + "\n"))
	conn.Close()
	// }

	// for {
	// 	netData, err := bufio.NewReader(conn).ReadString('\n')
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	if strings.TrimSpace(string(netData)) == "STOP" {
	// 		fmt.Println("Exiting TCP server!")
	// 		return
	// 	}
	// }

}

func main() {
	fmt.Println("Launching server...")
	fmt.Println("Listen on port")
	ln, err := net.Listen("tcp", "10.36.15.220:9091")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Accept connection on port")
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// defer conn.Close()
		fmt.Println("Calling handleConnection")
		go handleConnection(conn)
	}

	// for {
	// 	netData, err := bufio.NewReader(conn).ReadString('\n')
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	if strings.TrimSpace(string(netData)) == "STOP" {
	// 		fmt.Println("Exiting TCP server!")
	// 		return
	// 	}
	// }

}
