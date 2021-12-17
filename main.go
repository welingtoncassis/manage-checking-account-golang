package main

import (
	account "banco/accounts"
	"banco/customers"
	"fmt"
)

func main() {

	firstPerson := customers.Person{Name: "Welington", Phone: "9999999"}
	firstAccount := account.Account{Holder: firstPerson, NumberAgency: 589, NumberAccount: 123456}
	firstAccount.Deposit(125.5)

	secondPerson := customers.Person{Name: "Assis", Phone: "999999"}
	secondAccount := account.Account{Holder: secondPerson, NumberAgency: 589, NumberAccount: 123457}
	secondAccount.Deposit(10.5)

	withdrawalValue := 200.0
	fmt.Println(firstAccount.Withdraw(withdrawalValue))

	status, value := firstAccount.Deposit(100.0)
	fmt.Println(status, value)

	transferStatus := firstAccount.Transfer(50, &secondAccount)
	fmt.Println(transferStatus)

	fmt.Println(firstAccount)
	fmt.Println(secondAccount)

}
