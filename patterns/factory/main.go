package main

import (
	"fmt"
	"log"

	"factory/creation"
)

func main() {
	pmCash, err := creation.GetPaymentMethod(creation.Cash)
	if err != nil {
		log.Println(err)
		return
	}

	msg := pmCash.Pay(20)
	fmt.Println(msg)

	pmCredit, err := creation.GetPaymentMethod(creation.CreditCard)
	if err != nil {
		log.Println(err)
		return
	}

	msg = pmCredit.Pay(20)
	fmt.Println(msg)
}
