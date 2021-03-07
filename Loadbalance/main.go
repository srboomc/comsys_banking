package main

import (
	"flag"
	"io"
	"log"
	"net"
	"strings"
)

// Backends define servers
type Backends struct {
	servers []string
	num     int
}

// Choose the server to load
func (back *Backends) Choose() string {
	i := back.num % len(back.servers)
	back.num++
	return back.servers[i]
}

func (back *Backends) String() string {
	return strings.Join(back.servers, ", ")
}

var (
	frontendAddr = flag.String("frontendAddr", "127.0.0.1:7777", "Frontend Address") //143.198.196.98
	backendAddr  = flag.String("backendAddr", "localhost:3000", "Backend Address")
	backends     *Backends
)

func init() {
	flag.Parse()

	if *frontendAddr == "" {
		log.Printf("Enter Frontend Address")
	}

	servers := strings.Split(*backendAddr, ",")
	if len(servers) == 1 && servers[0] == "" {
		log.Printf("Enter Backend Address")
	}

	backends = &Backends{servers: servers}
}

func main() {
	listener, err := net.Listen("tcp", *frontendAddr)
	if err != nil {
		log.Fatalf("failed to setup listener %v", err)
	}

	log.Println("ReverseProxy Listening " + *frontendAddr)

	for {
		frontendConn, err := listener.Accept()
		if err != nil {
			log.Fatalf("failed to accept listener %v", err)
		}
		log.Print("Accepted frontendConn")
		go rvproxy(frontendConn, *backendAddr)
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
