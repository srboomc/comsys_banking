package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

// var srv Server
type clientData struct {
	DataTypeChoose string
	SenderAcc      int
	RecieverAcc    int
	Money          int
}

type responseData struct {
	DataTypeChoose string `json:"datatypechoose"`
	SenderAcc      int    `json:"senderacc"`
	RecieverAcc    int    `json:"recieveracc"`
	Money          int    `json:"money"`
}

// func connHandled(c chan int, j int) {
func connHandled(c chan int, j int, jdata []byte) {
	// func connHandled(j int, jdata []byte) {
	Addr := "localhost:9091"
	// Addr := "10.36.15.220:9091"
	start := time.Now()
	// Addr := "10.36.7.166:4000"
	// msg := "hello world " + strconv.Itoa(j) + "\n"
	conn, err := net.Dial("tcp", Addr)
	if err != nil {
		log.Println("Dial failed for address: " + Addr)
	}
	// conn.Write([]byte(msg))
	conn.Write(jdata)
	// conn.Write(jdata)
	// var m clientData
	// log.Print(jdata)
	// err := json.Unmarshal(jdata, &m)
	// if err != nil {
	// 	log.Print("no json found")
	// }
	// log.Print(m)
	log.Print("print somthing " + strconv.Itoa(j))
	defer conn.Close()
	elasped := time.Since(start)
	log.Print(elasped)
	c <- 0
}

func main() {
	// start := time.Now()
	dataType := [5]string{"transfer", "withdraw", "deposit", "buystock", "sellstock"}
	n := 6
	var b clientData
	c := make(chan int)
	// dataType[rand.Intn(5)]
	// log.Print("hello")
	for i := 0; i < n; i++ {
		v := dataType[rand.Intn(5)]
		// log.Print(v)
		m := clientData{
			DataTypeChoose: v,
			SenderAcc:      123455,
			RecieverAcc:    123456,
			Money:          200}
		if v != "transfer" {
			m.RecieverAcc = 0
		}
		log.Print(m)
		jsonDato, _ := json.Marshal(m)
		// if err != nil {
		// 	log.Println("json error")
		// }
		// os.Stdout.Write(jsonDato)
		// fmt.Println(jsonDato)
		// fmt.Println(string(jsonDato))
		err := json.Unmarshal(jsonDato, &b)
		if err != nil {
			log.Print("json not found")
		}
		// log.Print(b)
		// log.Print(jsonData)
		// jsonData == []byte()
		// log.Print(k)
		// log.Print(i)
		// go connHandled(c, i)
		go connHandled(c, i, jsonDato)
		// connHandled(i, jsonData)
	}
	// elasped := time.Since(start)
	// log.Print(elasped)
	<-c
}
