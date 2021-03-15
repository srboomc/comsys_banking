// https://github.com/tonymackay/go-yahoo-finance/blob/master/client.go

package main

import (
	"fmt"
	"log"
	"github.com/go-redis/redis"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// "reflect"
	"github.com/tonymackay/go-yahoo-finance"
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
var symbol string

var symbolName string
var currentPrice float64

func main() {

	db, err1 := sql.Open("mysql", "root:CIEBanking05comsys.@tcp(143.198.196.98:3306)/Banking")
	log.Print(err1)
	sqliteHandler.Conn = db
	fmt.Println("successfully connected to mysql")


	
	fmt.Printf("Get Keys : ")
	fmt.Scan(&symbol)

	client := newClient()

	err := ping(client)
	if err!= nil {
		fmt.Println(err)
	}

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
func set(client *redis.Client) error {

	info := client.Set(symbolName,currentPrice,0).Err()
	
	if info != nil {
		return info
		}
	return nil

	
}

func get(client *redis.Client) error {
		
	val, err := client.Get(symbol).Result()

	if err != nil {
		start := time.Now()
		getFromSQL(symbol)
		elasped := time.Since(start)
		log.Print(elasped)
	}else{
		start1 := time.Now()
		fmt.Println(symbol, val,"USD")
		elasped1 := time.Since(start1)
		log.Print(elasped1)
	}

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
			symbolName = ele.symbol
			currentPrice = ele.price
			fmt.Printf("%s  %f USD\n", symbolName, currentPrice)
			set(newClient())

			return (symbolName)
		}
	} else {
		var symbol string
	
		fmt.Printf("Symbol does not exist please re-type : ")
		fmt.Scan(&symbol)
		result, err := yahoo.Quote(symbol)
		if err != nil {
			log.Panic(err)
		}

		fmt.Printf("%v Price: %v %s\n",symbol,
			result.QuoteSummary.Result[0].Price.RegularMarketPrice.Raw,
			result.QuoteSummary.Result[0].Price.Currency,
		)
		usprice := (result.QuoteSummary.Result[0].Price.RegularMarketPrice.Raw)
		fmt.Print(usprice)
		insert, err := sqliteHandler.Conn.Query("INSERT INTO Stock(symbol, price) VALUES(?, ?)",symbol,usprice)
		
		if err != nil {
			panic(err.Error())
		}
		// be careful deferring Queries if you are using transactions
		defer insert.Close()
		
	}
	return ""

}