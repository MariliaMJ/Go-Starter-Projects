package accounts

import "bank/customers"

type SavingsAccount struct {
	AccountHolder                              customers.AccountHolder
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (c *SavingsAccount) Withdraw(withdrawalValue float64) string {
	canWithdraw := withdrawalValue > 0 && withdrawalValue <= c.balance
	if canWithdraw {
		c.balance -= withdrawalValue
		return "Withdrawal successfully executed"
	} else {
		return "Insufficient balance"
	}
}

func (c *SavingsAccount) Deposit(DepositValue float64) (string, float64) {
	if DepositValue > 0 {
		c.balance += DepositValue
		return "Deposi successfully executed", c.balance
	} else {
		return "Deposit value less than sero", c.balance
	}
}

func (c *SavingsAccount) Getbalance() float64 {
	return c.balance
}
