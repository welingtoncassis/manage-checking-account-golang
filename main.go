package main

import "fmt"

type Account struct {
	holder        string
	numberAgency  int
	numberAccount int
	balance       float64
}

func main() {

	firstAccount := Account{"Welington", 589, 123456, 125.5}

	// var firstAccount *Account
	// firstAccount = new(Account)
	// firstAccount.balance = 500

	withdrawalValue := 200.0
	fmt.Println(firstAccount.Withdraw(withdrawalValue))

}

func (c *Account) Withdraw(withdrawalValue float64) string {

	isOk := withdrawalValue > 0 && withdrawalValue <= c.balance
	if isOk {
		c.balance -= withdrawalValue
		return "successful withdrawal"
	} else {
		return "insufficient funds"
	}
}
