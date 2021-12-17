package account

import (
	"banco/customers"
)

type Account struct {
	Holder        customers.Person
	NumberAgency  int
	NumberAccount int
	balance       float64
}

func (a *Account) Withdraw(withdrawalValue float64) string {

	isOk := withdrawalValue > 0 && withdrawalValue <= a.balance
	if isOk {
		a.balance -= withdrawalValue
		return "successful withdrawal"
	} else {
		return "insufficient funds"
	}
}

func (a *Account) Deposit(amountdeposit float64) (string, float64) {
	isOK := amountdeposit > 0

	if isOK {
		a.balance += amountdeposit
		return "deposit done successfully", a.balance
	} else {
		return "error when making the deposit", a.balance
	}
}

func (a *Account) Transfer(transferValue float64, destinationAccount *Account) bool {
	isOK := transferValue <= a.balance && transferValue > 0

	if isOK {
		a.balance -= transferValue
		destinationAccount.Deposit(transferValue)
		return true
	} else {
		return false
	}
}

func (a *Account) GetBalance() float64 {
	return a.balance
}
