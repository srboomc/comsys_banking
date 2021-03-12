package main

import (
	"fmt"
	"log"

	"github.com/tonymackay/go-yahoo-finance"
)

func main() {
	var symbol string
	fmt.Printf("Symbol of stock : ")
	fmt.Scan(&symbol)
	result, err := yahoo.Quote(symbol)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%v Price: %v %s\n",symbol,
		result.QuoteSummary.Result[0].Price.RegularMarketPrice.Raw,
		result.QuoteSummary.Result[0].Price.Currency,
	)
	price := result.QuoteSummary.Result[0].Price.RegularMarketPrice.Raw
	fmt.Print(price)
}