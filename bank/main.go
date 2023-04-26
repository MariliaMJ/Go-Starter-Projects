package main

import (
	"fmt"
	"bank/accounts"
)

func PayBill(bill verifyBill, billValue float64) {
	bill.Withdraw(billValue)
}

type verifyBill interface {
	Withdraw(valor float64) string
}

func main() {
	contaDoDenis := accounts.SavingsAccount{}
	contaDoDenis.Deposit(100)
	PayBill(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.Getbalance())

	contaDaLuisa := accounts.CheckingAccount{}
	contaDaLuisa.Deposit(500)
	PayBill(&contaDaLuisa, 1000)

	fmt.Println(contaDaLuisa.GetBalance())

}
