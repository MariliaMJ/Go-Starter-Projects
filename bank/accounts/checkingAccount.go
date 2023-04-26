package accounts

import "bank/customers"

type CheckingAccount struct {
	Titular                    customers.AccountHolder
	AgencyNumber, AccountNumber int
	balance                      float64
}

func (c *CheckingAccount) Withdraw(withdrawalValue float64) string {
	canWithdraw := withdrawalValue > 0 && withdrawalValue <= c.balance
	if canWithdraw {
		c.balance -= withdrawalValue
		return "Withdrawal successfully executed"
	} else {
		return "Insufficient balance"
	}
}

func (c *CheckingAccount) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		c.balance += depositValue
		return "Deposit successfully executed", c.balance
	} else {
		return "Deposit value less than zero", c.balance
	}
}

func (c *CheckingAccount) Transfer(transferValue float64, destinationAccount *CheckingAccount) bool {
	if transferValue < c.balance && transferValue > 0 {
		c.balance -= transferValue
		destinationAccount.Deposit(transferValue)
		return true
	} else {
		return false
	}
}

func (c *CheckingAccount) GetBalance() float64 {
	return c.balance
}