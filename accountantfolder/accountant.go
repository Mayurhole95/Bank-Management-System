package acct

import (
	"fmt"
	"math/rand"
)

type customer struct {
	email    string
	password int
	phone    int
	acc_no   int
	balance  int
}

var map1 = make(map[int]customer)
var acc_no = 1000

func Verify(acc_no int, password int) bool {
	temp := map1[acc_no]
	//	fmt.Println(password)

	if password == temp.password {

		return true
	} else {
		return false
	}

}
func Balance(acc_no int) int {
	temp := map1[acc_no]
	return temp.balance
}

func Depositupdatebalance(acc_no, amount int) int {

	var temp customer = map1[acc_no]
	temp.balance += amount
	map1[acc_no] = temp
	return temp.balance

}

func WithdrawUpdatebalance(acc_no, amount int) int {
	var temp customer = map1[acc_no]
	temp.balance -= amount
	map1[acc_no] = temp
	return temp.balance
}

func create_account(c1 customer) {

	var email string
	var phone int

	fmt.Println("Enter your e-mail: ")
	fmt.Scanln(&email)
	c1.email = email
	fmt.Println("Enter phone no: ")
	fmt.Scanln(&phone)
	c1.phone = phone

	fmt.Println("Account Created Successfully !")
	fmt.Println("Your Account Number is: ", acc_no)
	c1.acc_no = acc_no

	password := rand.Intn(10000)
	fmt.Println("Your Password is: ", password)
	c1.password = password

	map1[acc_no] = c1
	acc_no = acc_no + 1
}

func delete_account() {
	var acc_no int
	fmt.Println("Enter account number to delete account : ")
	fmt.Scanln(&acc_no)

	delete(map1, acc_no)

	fmt.Println("Account Deleted Successfully")
}

func Accountant() {
	var input int
	var email string
	var password string
	fmt.Println("Enter your Accountant E-mail Id: ")
	fmt.Scanln(&email)
	fmt.Println("Enter your password: ")
	fmt.Scanln(&password)

	if email == "account@bank.com" && password == "josh@123" {
		fmt.Println("Login Successful")
	} else {
		fmt.Println("Invalid credentials, Try again ")
		return
	}
	for {
		fmt.Println("1. Create Account")
		fmt.Println("2. Delete Account")
		fmt.Println("3. Exit")
		fmt.Scanln(&input)
		flag := false

		switch input {
		case 1:
			var c1 customer
			create_account(c1)
		case 2:
			delete_account()
		case 3:
			flag = true

		}
		if flag {
			break
		}
	}

}
