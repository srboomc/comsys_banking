package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// var srv Server
type clientData struct {
	dataTypeChoose string
	senderAcc      int64
	recieverAcc    int64
	money          int64
}

// func connHandled(c chan int, j int, jdata []byte) {
func connHandled(j int, jdata []byte) {
	// Addr := "localhost:7777"
	// Addr := "10.36.15.220:9091"
	// // msg := "hello world\n"
	// conn, err := net.Dial("tcp", Addr)
	// if err != nil {
	// 	log.Println("Dial failed for address: " + Addr)
	// }
	// conn.Write([]byte(msg))
	// conn.Write(jdata)
	var m clientData
	log.Print(jdata)
	err := json.Unmarshal(jdata, &m)
	if err != nil {
		log.Print("no json found")
	}
	log.Print(m)
	log.Print("print somthing " + strconv.Itoa(j))
	// defer conn.Close()
	// c <- 0
}

func main() {
	start := time.Now()
	dataType := [5]string{"transfer", "withdraw", "deposit", "buystock", "sellstock"}
	n := 5
	var b clientData
	// c := make(chan int)
	// dataType[rand.Intn(5)]
	// log.Print("hello")
	for i := 0; i < n; i++ {
		v := dataType[rand.Intn(5)]
		// log.Print(v)
		m := &clientData{v, 123455, 123456, 200}
		if v != "transfer" {
			m.recieverAcc = 0
		}
		log.Print(m)
		jsonData, err := json.Marshal(m)
		if err != nil {
			log.Println("json error")
		}
		fmt.Println(string(jsonData))
		err = json.Unmarshal(jsonData, &b)
		if err != nil {
			log.Print("json not found")
		}
		// log.Print(b)
		log.Print(jsonData)
		// jsonData == []byte()
		// log.Print(k)
		// log.Print(i)
		// go connHandled(c, i, jsonData)
		// connHandled(i, jsonData)
	}
	elasped := time.Since(start)
	log.Print(elasped)
	// <-c
}
