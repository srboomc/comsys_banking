package main

import (
	"database/sql"
	"fmt"

	// "net/http"

	_ "github.com/go-sql-driver/mysql"
)

type SQLHandler struct {
	Conn *sql.DB
}

type accountInfo struct {
	ID            int    `json:"id"`
	Date          string `json:"date"`
	AccountName   string `json:"accountName"`
	AccountNumber int    `json:"accountNumber"`
	Phone         string `json:"phone"`
	Balance       int    `json:"balance"`
}

type balance struct {
	AccountName string `json:"accountName"`
	Balance     int    `json:"balance"`
}

var sqliteHandler SQLHandler

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)

	}
}

func checkBalance() {
	data, err := sqliteHandler.Conn.Query("select Account_Name, Balance from Account_info")

	var result []balance

	for data.Next() {
		var res balance
		err = data.Scan(&res.AccountName, &res.Balance)
		checkErr(err)
		result = append(result, res)
	}

	if len(result) != 0 {
		for _, ele := range result {
			fmt.Printf("Account : %s ---> Balance: %d$ \n", ele.AccountName, ele.Balance)

		}
	} else {
		fmt.Println("Data not found!!")
	}
}

func retrieveUserInfo() {
	data, err := sqliteHandler.Conn.Query("select id, date, Account_Name, Account_Number, Phone, Balance from Account_info")

	var result []accountInfo

	for data.Next() {
		var res accountInfo
		err = data.Scan(&res.ID, &res.Date, &res.AccountName, &res.AccountNumber, &res.Phone, &res.Balance)
		checkErr(err)
		result = append(result, res)
	}
	if len(result) != 0 {
		for _, ele := range result {
			fmt.Printf("Account Id: %d\n", ele.ID)
			fmt.Printf("Name: %s\n", ele.AccountName)
			fmt.Printf("Account Number: %d\n", ele.AccountNumber)
			fmt.Printf("Date: %s\n", ele.Date)
			fmt.Printf("Phone: %d\n", ele.Phone)
			fmt.Printf("Balance: %d\n", ele.Balance)
		}
	} else {
		fmt.Println("Data not found!!!")
	}
}

func main() {
	fmt.Println("open mysql")

	db, err := sql.Open("mysql", "root:CIEBanking05comsys.@tcp(143.198.196.98:3306)/Banking")
	checkErr(err)

	sqliteHandler.Conn = db

	defer db.Close()
	// retrieveUserInfo()
	checkBalance()
	fmt.Println("successfully connected to mysql")

}
