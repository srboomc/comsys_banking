package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"sync"

	// "net/http"

	_ "github.com/go-sql-driver/mysql"
)

type SQLHandler struct {
	Conn *sql.DB
}

type accountInfo struct {
	Date          string `json:"date"`
	AccountName   string `json:"accountName"`
	AccountNumber int    `json:"accountNumber"`
	Phone         string `json:"phone"`
	Balance       int    `json:"balance"`
}

type balance struct {
	// AccountNumber string `json:"accountNumber"`
	Balance int `json:"balance"`
}

var sqliteHandler SQLHandler

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)

	}
}

var m sync.Mutex

func tranfer(sndAccount int, recvAccount int, amount int) string {
	m.Lock()
	sndBalance := retrieveUserBalance(strconv.Itoa(sndAccount))
	if amount > sndBalance {
		return ("not enough money!!!")
	}
	remain := retrieveUserBalance(strconv.Itoa(recvAccount))
	sqliteHandler.Conn.Query("SET SQL_SAFE_UPDATES = 0")

	sqliteHandler.Conn.Query("insert into Transaction_Table (timestamp, senderAccount, receiverAccount, amount) values ( NOW()," + strconv.Itoa(sndAccount) + "," + strconv.Itoa(recvAccount) + "," + strconv.Itoa(amount) + ")")
	sqliteHandler.Conn.Query("update Account_info set Balance = " + strconv.Itoa(sndBalance-amount) + " where Account_Number = " + strconv.Itoa(sndAccount))
	sqliteHandler.Conn.Query("update Account_info set Balance = " + strconv.Itoa(remain+amount) + " where Account_Number = " + strconv.Itoa(recvAccount))
	m.Unlock()
	return ("done")
}

func withdraw(AccountNumber int, amount int) int {
	balance := retrieveUserBalance(strconv.Itoa(AccountNumber))

	sqliteHandler.Conn.Query("update Account_info set Balance = " + strconv.Itoa(balance-amount) + " where Account_Number = " + strconv.Itoa(AccountNumber))

	remain := retrieveUserBalance(strconv.Itoa(AccountNumber))

	return remain
}

func deposit(AccountNumber int, amount int) int {
	balance := retrieveUserBalance(strconv.Itoa(AccountNumber))

	sqliteHandler.Conn.Query("update Account_info set Balance = " + strconv.Itoa(balance+amount) + " where Account_Number = " + strconv.Itoa(AccountNumber))

	remain := retrieveUserBalance(strconv.Itoa(AccountNumber))

	return remain
}

func retrieveUserBalance(AccountNumber string) int {

	data, err := sqliteHandler.Conn.Query("select Balance from Account_info where Account_Number =" + AccountNumber)

	var result []balance

	for data.Next() {
		var res balance
		err = data.Scan(&res.Balance)
		checkErr(err)
		result = append(result, res)
	}
	if len(result) != 0 {
		for _, ele := range result {
			// fmt.Println(ele.Balance)
			return ele.Balance
		}
	} else {
		fmt.Println("Data not found!!!")
	}
	return 0
}

func checkBalance(AccountNumber int) {
	data, err := sqliteHandler.Conn.Query("select date, Account_Name, Account_Number, Phone, Balance from Account_info where Account_Number = " + strconv.Itoa(AccountNumber))

	var result []accountInfo

	for data.Next() {
		var res accountInfo
		err = data.Scan(&res.Date, &res.AccountName, &res.AccountNumber, &res.Phone, &res.Balance)
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

func main() {
	fmt.Println("open mysql")

	db, err := sql.Open("mysql", "root:CIEBanking05comsys.@tcp(143.198.196.98:3306)/Banking")
	checkErr(err)

	sqliteHandler.Conn = db

	defer db.Close()
	fmt.Println("successfully connected to mysql")
	// checkBalance(123455678)
	// fmt.Println(tranfer(98765432, 123455678, 1000))

}
