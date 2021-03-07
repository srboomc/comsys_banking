package main

import (
	"io"
	"log"
	"net"
)

func main() {
	frontendAddr := "localhost:7777"
	backendAddr := "localhost:1234"

	listener, err := net.Listen("tcp", frontendAddr)
	if err != nil {
		log.Fatalf("failed to setup listener %v", err)
	}
	log.Println("ReverseProxy Listening on " + frontendAddr)
	for {
		frontendConn, err := listener.Accept()
		if err != nil {
			log.Fatalf("failed to accept listener %v", err)
		}
		log.Print("Accepted frontendConn")

		go rvproxy(frontendConn, backendAddr)
		log.Print("Done!!!")
	}
}

func rvproxy(frontendConn net.Conn, backendAddr string) {
	backendConn, err := net.Dial("tcp", backendAddr)
	if err != nil {
		log.Fatalf("Dial failed for address" + backendAddr)

	}
	defer frontendConn.Close()
	log.Print("send request")
	io.Copy(backendConn, frontendConn)
	// netData, err := bufio.NewReader(backendConn).ReadString('\n')
	// log.Print("dataSend")
	io.Copy(frontendConn, backendConn)
}

//reference: https://github.com/darkhelmet/balance/blob/master/tcp.go
