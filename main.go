package main

import (
	"fmt"
	acn "newbank/accountantfolder"
	acm "newbank/customerfolder"
)

var acc_no = 1

func main() {

	var choice int

	for {
		flag := false
		fmt.Println("----------Welcome--------")
		fmt.Println("Press 1 for Accountant Login")
		fmt.Println("Press 2 for Customer Login")
		fmt.Println("Press 3 to Exit")
		fmt.Println("Enter your Choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			acn.Accountant()

		case 2:
			acm.Customer()

		case 3:

			flag = true
		default:
			fmt.Println("Wrong Input")

		}
		if flag {
			break
		}

	}
}
