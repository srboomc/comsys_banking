package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var listAccount []accountDetail

func main() {
	var password string
	fmt.Println("Password for admin is admin")
	fmt.Println("Password for user is user")
	fmt.Printf("Enter your password : ")
	fmt.Scan(&password)
	if password == "user" {
		fmt.Printf("Password match!\nLoading")
		for i := 0; i < 5; i++ {
			// time.Sleep(1 * time.Second)
			fmt.Printf(".")
		}
		fmt.Println("")
		mainMenuUser()
	} else if password == "admin" {
		fmt.Printf("Password match!\nLoading")
		for i := 0; i < 5; i++ {
			time.Sleep(1 / 2 * time.Second)
			fmt.Printf(".")
		}
		fmt.Println("")
		mainMenuAdmin()
	} else {
		exitProgram()
	}
}

func mainMenuUser() {
	var selection int
	fmt.Println("-------- User Main Menu -----------")
	fmt.Printf("1.Create Account\n")
	fmt.Printf("2.Transaction\n")
	fmt.Printf("3.Check Balance\n")
	fmt.Println("4.Log out")
	fmt.Printf("Exit program, type 0\n")
	fmt.Printf("Enter your selection :")
	fmt.Scan(&selection)

	switch {
	case selection == 1:
		createAccount()
	case selection == 2:
		transaction()
	case selection == 3:
		balance()
	case selection == 4:
		main()
	case selection == 0:
		exitProgram()
	default:
		fmt.Println("Invalid")

	}
}

func mainMenuAdmin() {
	var selection int

	fmt.Println("-------- Admin Main Menu -----------")
	fmt.Println("1. View all list of users' account")
	fmt.Println("Type 0 to exit the admin menu")
	fmt.Printf("Enter your choice :")
	fmt.Scan(&selection)
	switch {
	case selection == 1:
		showAccount()
	case selection == 0:
		main()
	}
}

type accountDetail struct {
	date           string
	Account_Name   string
	Account_Number int
	Phone          string
	balance        int
}

func createAccount() {
	var accountName string
	var phone string
	var balance int
	min := 100000000000
	max := 999999999999
	randomNumber := rand.Intn(max-min) + min

	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02")

	fmt.Printf("Account name: ")
	fmt.Scan(&accountName)
	fmt.Printf("Phone Number: ")
	fmt.Scan(&phone)
	fmt.Printf("Deposit: ")
	fmt.Scan(&balance)
	fmt.Println("")

	a1 := (accountDetail{currentDate,
		accountName,
		randomNumber,
		phone,
		balance})

	listAccount = append(listAccount, a1)
	fmt.Println("SAVED!")
	mainMenuUser()
}

func transaction() {
	var choice int

	fmt.Printf("-----------------Transaction----------------\n")
	fmt.Println("1. Deposit")
	fmt.Println("2. Withdraw")
	fmt.Println("3. Transfer money")
	fmt.Printf("Enter your choice: ")
	fmt.Scan(&choice)
	switch {
	case choice == 1:
		deposit()
	case choice == 2:
		withdraw()
	}
}

func deposit() {
	var number int
	var money int

	fmt.Printf("Write your account number : ")
	fmt.Scan(&number)
	fmt.Printf("Amount of money : ")
	fmt.Scan(&money)
	for i := 0; i < len(listAccount); i++ {
		if number == listAccount[i].Account_Number {
			listAccount[i].balance += money
			fmt.Println("Deposit success!")
		} else {
			fmt.Println("Invalid")
		}
		mainMenuUser()
	}
}

func withdraw() {
	var number int
	var money int

	fmt.Printf("Write your account number : ")
	fmt.Scan(&number)
	fmt.Printf("Amount of money : ")
	fmt.Scan(&money)
	for i := 0; i < len(listAccount); i++ {
		if number == listAccount[i].Account_Number {
			if listAccount[i].Account_Number >= money {
				listAccount[i].balance -= money
			} else {
				fmt.Println("Not enough money!")
			}
		}
		mainMenuUser()
	}
}

func balance() {
	var number int

	fmt.Printf("Type your Account number : \n")
	fmt.Scan(&number)

	for i := 0; i < len(listAccount); i++ {
		if number == listAccount[i].Account_Number {
			fmt.Println("Your balance is : ", listAccount[i].balance)
		}
	}
	mainMenuUser()

}

func showAccount() {
	if len(listAccount) == 0 {
		fmt.Println("There is no account")
	} else {
		fmt.Println("Date \t\t Name \t\t Account Number \t\t Phone \t\t\tBalance")
		for i := 0; i < len(listAccount); i++ {
			fmt.Println(listAccount[i].date, "\t", listAccount[i].Account_Name, "\t\t", listAccount[i].Account_Number, "\t\t\t", listAccount[i].Phone, "\t\t", listAccount[i].balance)

		}
		time.Sleep(1 * time.Second)
		mainMenuAdmin()
	}
}
func exitProgram() {
	os.Exit(3)
}
