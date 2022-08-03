package cust

import (
	"fmt"
	all "newbank/accountantfolder"
	"time"
)

type hist struct {
	srno               int
	date               string
	transaction_type   string
	transaction_amount int
	total_bal          int
}

var withdraw_amount int
var deposit_amount int
var newbala int

var historymap = map[int][]hist{}

func Withdraw(acc_no int, withdraw_amount int) int {
	wd := all.WithdrawUpdatebalance(acc_no, withdraw_amount)
	_, n := historymap[acc_no]
	d := time.Now()
	if n {
		var slice []hist = historymap[acc_no]
		cnt := 0
		for i := 0; i < len(slice); i++ {
			cnt++
		}

		var h2 hist = slice[cnt-1]
		sr := h2.srno

		var h1 hist
		h1.srno = sr + 1
		h1.date = d.Format("01-02-2006")
		h1.transaction_type = "Withdraw"
		h1.transaction_amount = withdraw_amount
		h1.total_bal = all.Balance(acc_no)
		slice = append(slice, h1)
		historymap[acc_no] = slice

	} else {
		var empty []hist
		var h1 hist
		h1.srno = 1
		h1.date = d.Format("01-02-2006")
		h1.transaction_type = "Withdraw"
		h1.transaction_amount = withdraw_amount
		h1.total_bal = all.Balance(acc_no)
		empty = append(empty, h1)
		historymap[acc_no] = empty
	}
	return wd

}
func Deposit(acc_no int, deposit_amount int) int {
	dd := all.Depositupdatebalance(acc_no, deposit_amount)
	_, n := historymap[acc_no]
	d := time.Now()
	if n {
		var slice []hist = historymap[acc_no]
		cnt := 0
		for i := 0; i < len(slice); i++ {
			cnt++
		}

		var h2 hist = slice[cnt-1]
		sr := h2.srno

		var h1 hist
		h1.srno = sr + 1
		h1.date = d.Format("01-02-2006")
		h1.transaction_type = "Deposit"
		h1.transaction_amount = deposit_amount
		h1.total_bal = all.Balance(acc_no)
		slice = append(slice, h1)
		historymap[acc_no] = slice

	} else {
		var empty []hist
		var h1 hist
		h1.srno = 1
		h1.date = d.Format("01-02-2006")
		h1.transaction_type = "Deposit"
		h1.transaction_amount = deposit_amount
		h1.total_bal = all.Balance(acc_no)
		empty = append(empty, h1)
		historymap[acc_no] = empty
	}
	return dd
}

func trans_hist(acc_no int) {
	var ma []hist = historymap[acc_no]
	for i := 0; i < len(ma); i++ {
		var temp hist = ma[i]
		fmt.Println("-------Transaction History-------")
		fmt.Println("Sr. No.: ", temp.srno)
		fmt.Println("Date : ", temp.date)
		fmt.Println("Type Of Transaction :", temp.transaction_type)
		fmt.Println("Transaction Amount : ", temp.transaction_amount)
		fmt.Println("Total Balance : ", temp.total_bal)
	}
}

func Customer() {
	var ch int
	balancee := 0
	var password int
	var acc_no int
	flag := false

	fmt.Println("Enter your account Id: ")
	fmt.Scanln(&acc_no)

	fmt.Println("Enter your password")
	fmt.Scanln(&password)
	//fmt.Println(all.Verify(acc_no, password))
	if all.Verify(acc_no, password) {
		fmt.Println("Logged in ")
		for {
			fmt.Println("1. View Balance")
			fmt.Println("2. Withdraw")
			fmt.Println("3. Deposit")
			fmt.Println("4. Transaction History")
			fmt.Println("5. Exit")
			fmt.Println("Enter your choice: ")
			fmt.Scanln(&ch)

			switch ch {
			case 1:

				balancee = all.Balance(acc_no)
				fmt.Println("Balance : ", balancee)
			case 2:
				fmt.Println("Enter amount to withdraw: ")
				fmt.Scanln(&withdraw_amount)

				newbal := Withdraw(acc_no, withdraw_amount)
				fmt.Println("New balance: ", newbal)
			case 3:
				fmt.Println("Enter amount to deposit: ")
				fmt.Scanln(&deposit_amount)
				newbala = Deposit(acc_no, deposit_amount)
				fmt.Println("New balance: ", newbala)
			case 4:
				//code for history
				trans_hist(acc_no)

			case 5:
				flag = true
			}
			if flag {
				break
			}
		}
	}

}
