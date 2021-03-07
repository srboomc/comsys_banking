package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func ServerChecker(tt time.Time) {
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatalf("Dial failed for address localhost:3000")
		return
	}
	select {
	case <-time.After(time.Second * 1):
		fmt.Println(time.Now(), "Healthy")
	}
	defer conn.Close()
}

func main() {
	timecheck(20*time.Millisecond, ServerChecker)
}

func timecheck(timedu time.Duration, fun func(time.Time)) {
	for val := range time.Tick(timedu) {
		ServerChecker(val)
	}
}
