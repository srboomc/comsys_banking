package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var listAccount []accountDetail
var phone string

func main() {
	var selection int
	fmt.Println("1.Create Account")
	fmt.Println("2.Login")
	fmt.Print("Enter your seletion : ")
	fmt.Scan(&selection)
	switch {
	case selection == 1:
		createAccount()
	case selection == 2:
		loginpage()
	}
}

type accountDetail struct {
	Date           string
	Account_Name   string
	Account_Number int
	Phone          string
	Password       string
	Balance        uint32
}

func createAccount() {
	var accountName string
	var phone string
	var password string
	var deposit uint32

	min := 100000000000
	max := 999999999999
	randomNumber := rand.Intn(max-min) + min

	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02")

	fmt.Printf("Account name: ")
	fmt.Scan(&accountName)
	fmt.Printf("Phone Number: ")
	fmt.Scan(&phone)
	fmt.Printf("Password: ")
	fmt.Scan(&password)
	fmt.Printf("Deposit: ")
	fmt.Scan(&deposit)

	a1 := (accountDetail{currentDate,
		accountName,
		randomNumber,
		phone,
		password,
		deposit})
	b, err := json.Marshal(a1)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	listAccount = append(listAccount, a1)

	loginpage()
}

func loginpage() {
	var password string
	fmt.Printf("\nPhone no : ")
	fmt.Scan(&phone)
	fmt.Printf("Password : ")
	fmt.Scan(&password)

	for i := 0; i < len(listAccount); i++ {
		if listAccount[i].Phone == phone && listAccount[i].Password == password {
			transaction()
		} else {
			fmt.Println("Invalid password")
		}
	}
}

func transaction() {
	var selection int
	fmt.Println("\n -------- User transaction Menu -----------")
	fmt.Println("1.Deposit money")
	fmt.Println("2.Withdraw money")
	fmt.Println("3.Transfer money")
	fmt.Println("4.Check balance")
	fmt.Println("5.Logout")
	fmt.Println("Exit program, type 0")
	fmt.Printf("Enter your selection :")
	fmt.Scan(&selection)

	switch {
	case selection == 1:
		deposit()
	case selection == 2:
		withdraw()
	case selection == 3:
		fmt.Println("No service yet")
	case selection == 4:
		balance()
	case selection == 5:
		main()
	case selection == 0:
		exitProgram()
	default:
		fmt.Println("Invalid")

	}
}

func deposit() {
	var money uint32

	fmt.Printf("Amount of money : ")
	fmt.Scan(&money)
	for i := 0; i < len(listAccount); i++ {
		if phone == listAccount[i].Phone {
			listAccount[i].Balance += money
			fmt.Println("Deposit success!")
		} else {
			fmt.Println("Invalid")
		}
		transaction()
	}
}

func withdraw() {
	var money uint32

	fmt.Printf("Amount of money : ")
	fmt.Scan(&money)
	for i := 0; i < len(listAccount); i++ {
		if phone == listAccount[i].Phone {
			if listAccount[i].Balance >= money {
				listAccount[i].Balance -= money
			} else {
				fmt.Println("Not enough money!")
			}
		}
		transaction()
	}
}

func balance() {
	var number string
	var choice int

	for i := 0; i < len(listAccount); i++ {
		if phone == listAccount[i].Phone {

			fmt.Println("Your balance is : ", listAccount[i].Balance)
		}
	}
	fmt.Printf("type 0 to exit : ")
	fmt.Scan(&number)
	if choice == 0 {
		transaction()
	}

}

func showAccount() {
	var choice int

	if len(listAccount) == 0 {
		fmt.Println("There is no account")
	} else {
		fmt.Println("Date \t\t Name \t\t Account Number \t\t Phone \t\t\tBalance")
		for i := 0; i < len(listAccount); i++ {
			fmt.Println(listAccount[i].Date, "\t", listAccount[i].Account_Name, "\t\t", listAccount[i].Account_Number, "\t\t\t", listAccount[i].Phone, "\t\t", listAccount[i].Balance)

		}
		time.Sleep(1 * time.Second)

	}
	fmt.Printf("Type 0 to exit : ")
	fmt.Scan(&choice)
	if choice == 0 {
		transaction()
	}
}

func exitProgram() {
	os.Exit(3)
}
