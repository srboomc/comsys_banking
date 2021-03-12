package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

// Backends define servers
type Backends struct {
	servers []string
	n       int
}

// Choose the server to load (Round Robin)
func (b *Backends) Choose() string {
	idx := b.n % len(b.servers)
	b.n++
	return b.servers[idx]
}

func (b *Backends) String() string {
	return strings.Join(b.servers, ", ")
}

var (
	bind = flag.String("bind", "10.36.7.166:4000", "The address to bind on") //143.198.196.98 #127.0.0.1
	// balance = flag.String("balance", "10.36.7.166:4001", "The backend servers to balance connections across, separated by commas")
	balance  = flag.String("balance", "127.0.0.1:4001,127.0.0.1:4002", "The backend servers to balance connections across, separated by commas")
	backends *Backends
)

func init() {
	flag.Parse()

	if *bind == "" {
		log.Fatalln("specify the address to listen on with -bind")
	}

	servers := strings.Split(*balance, ",")
	if len(servers) == 1 && servers[0] == "" {
		log.Fatalln("please specify backend servers with -backends")
	}

	backends = &Backends{servers: servers}
}

func copy(wc io.WriteCloser, r io.Reader) {
	defer wc.Close()
	io.Copy(wc, r)

}

func handleConnection(us net.Conn, server string) {
	ds, err := net.Dial("tcp", server)
	if err != nil {
		us.Close()
		log.Printf("failed to dial %s: %s", server, err)
		return
	}

	go copy(ds, us)
	go copy(us, ds)
}

func ServerChecker(tt time.Duration) {
	// conn, err := net.Dial("tcp", "localhost:3000")
	servers := strings.Split(*balance, ",")
	for {
		// if len(servers) == 1 && servers[0] == "" {
		// log.Fatalln("please specify backend servers with -backends")
		conn1, err := net.Dial("tcp", servers[0])
		if err != nil {
			// log.Fatalf("Dial failed for address localhost:3000")
			log.Fatalf("Dial failed for address %s", err)
			return
		}

		conn2, err := net.Dial("tcp", servers[1])
		if err != nil {
			// log.Fatalf("Dial failed for address localhost:3000")
			log.Fatalf("Dial failed for address %s", err)
			return
		}
		select {
		case <-time.After(tt):
			fmt.Println(time.Now(), "Healthy")
		}
		// defer conn.Close()
		defer conn1.Close()
		defer conn2.Close()
	}

	// conn, err := net.Dial("tcp", *balance)
	// if err != nil {
	// 	// log.Fatalf("Dial failed for address localhost:3000")
	// 	log.Fatalf("Dial failed for address %s", err)
	// 	return
	// conn, err := net.Dial("tcp", *balance)
	// if err != nil {
	// 	// log.Fatalf("Dial failed for address localhost:3000")
	// 	log.Fatalf("Dial failed for address %s", err)
	// 	return
	//tong nee perd wai pid sa
	// select {
	// case <-time.After(10 * time.Millisecond):
	// 	fmt.Println(time.Now(), "Healthy")
	// }
}

// func main() {
// 	timecheck(20*time.Millisecond, ServerChecker)
// }

func timecheck(timedu time.Duration, fun func(time.Duration)) {
	// for val := range time.Tick(timedu) {
	// 	ServerChecker(val)
	// }
	ServerChecker(timedu)
}

func main() {
	log.Printf("%s", *balance)
	go timecheck(5*time.Second, ServerChecker)
	ln, err := net.Listen("tcp", *bind)
	if err != nil {
		log.Fatalf("failed to bind: %s", err)
	}

	log.Printf("listening on %s, balancing %s", *bind, backends)
	count := 0
	for {
		// defer ln.Close()
		conn, err := ln.Accept()
		count++

		fmt.Println(count)
		if err != nil {
			log.Printf("failed to accept: %s", err)
			continue
		}
		go handleConnection(conn, backends.Choose())
		// count++
	}
	// fmt.Println(count)
}
