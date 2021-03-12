package main

import (
	"fmt"
	"log"
	"github.com/go-redis/redis"
	"github.com/tonymackay/go-yahoo-finance"
)

func main() {

	client := newClient()

	err := ping(client)
	if err != nil {
		fmt.Println(err)
	}

	err = set(client)
	if err != nil {
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
	var symbol string
	fmt.Printf("Symbol of stock : ")
	fmt.Scan(&symbol)
	
	result, err := yahoo.Quote(symbol)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%v Price: %v %s\n",symbol,
		result.QuoteSummary.Result[0].Price.RegularMarketPrice.Raw,
		result.QuoteSummary.Result[0].Price.Currency,)
	
	price := result.QuoteSummary.Result[0].Price.RegularMarketPrice.Raw

	info := client.Set(symbol,price,0).Err()
	if info != nil {
		return info
		}
	return nil
	
}

func get(client *redis.Client) error {
	val, err := client.Get("key2").Result()
	if err != nil {
		return (err)
	}
	fmt.Println("key2", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist

	return nil
}

