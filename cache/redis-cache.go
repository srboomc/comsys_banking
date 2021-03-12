// https://github.com/tonymackay/go-yahoo-finance/blob/master/client.go

package main

import (
	"fmt"
	"log"
	"github.com/go-redis/redis"
	// "github.com/tonymackay/go-yahoo-finance"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type SQLHandler struct {
	Conn *sql.DB
}

type stockdata struct {
	symbol string `json:"symbol"`
	price float64 `json:"price"`
}

var sqliteHandler SQLHandler
var err1 error

func main() {

	db, err1 := sql.Open("mysql", "root:CIEBanking05comsys.@tcp(143.198.196.98:3306)/Banking")
	log.Print(err1)
	sqliteHandler.Conn = db
	fmt.Println("successfully connected to mysql")


	client := newClient()

	err := ping(client)
	if err!= nil {
		fmt.Println(err)
	}

	// err = set(client)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	err = get(client)
	if err != nil {
		fmt.Println(err)
	}
	

}

func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}

// ping tests connectivity for redis (PONG should be returned)
func ping(client *redis.Client) error {
	pong, err := client.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println(pong, err)
	// Output: PONG <nil>

	return nil
}

// set executes the redis Set command
// func set(client *redis.Client) error {
// 	var symbol string
// 	fmt.Printf("Symbol of stock : ")
// 	fmt.Scan(&symbol)

// 	start := time.Now()

// 	result, err := yahoo.Quote(symbol)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	fmt.Printf("%v Price: %v %s\n",symbol,
// 		result.QuoteSummary.Result[0].Price.RegularMarketPrice.Raw,
// 		result.QuoteSummary.Result[0].Price.Currency,)
	
// 	price := result.QuoteSummary.Result[0].Price.RegularMarketPrice.Raw

// 	info := client.Set(symbol,price,0).Err()
// 	if info != nil {
// 		return info
// 		}
// 	elasped := time.Since(start)
// 	log.Print(elasped)

// 	return nil

	
// }

func get(client *redis.Client) error {
	var symbol string
	fmt.Printf("Get Keys : ")
	fmt.Scan(&symbol)

	start := time.Now()
	val, err := client.Get(symbol).Result()

	fmt.Println()
	if err != nil {
		return (err)
	}
	fmt.Println(symbol, val)

	elasped := time.Since(start)
	log.Print(elasped)
	

	start1 := time.Now()
	getFromSQL(symbol)
	elasped1 := time.Since(start1)
	log.Print(elasped1)

	return nil
}

func getFromSQL(symbol string)string{
	fmt.Println(symbol)
	data, err := sqliteHandler.Conn.Query("SELECT symbol, price FROM Stock WHERE symbol = '"+symbol+"'")
	
	var result []stockdata
	for data.Next() {
		var res stockdata
		err = data.Scan(&res.symbol, &res.price)
		result = append(result, res)
		fmt.Print(err)
	}

	if len(result) != 0 {
		for _ ,ele := range result {
			fmt.Printf("%s with price %f\n", ele.symbol, ele.price)
			return ele.symbol
		}
	} else {
		fmt.Println("data not found")
		return ("not found")
	}
	return ""

}

func insertToSQL(key string, value string){
	sqliteHandler.Conn.Query("INSERT INTO Stock (symbol,price) VALUES("+key+","+value+")")
	log.Print(key,value)
}
