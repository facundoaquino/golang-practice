package main

import (
	"flag"
	"fmt"
)

type PaymentStrategy interface {
	Pay()
}

type VisaStrategy struct{}

func (vs VisaStrategy) Pay() {
	fmt.Println("Pay with visa...")
}

type MasterStrategy struct{}

func (vs MasterStrategy) Pay() {
	fmt.Println("Pay with master...")
}

type DefaultStrategy struct{}

func (vs DefaultStrategy) Pay() {
	fmt.Println("Pay with default method here...")
}

func ChoiseStrategyToPay(paymehod string) PaymentStrategy {
	switch paymehod {
	case "visa":
		return VisaStrategy{}

	case "master":
		return MasterStrategy{}

	default:
		return DefaultStrategy{}
	}

}

func main() {

	payMethod := flag.String("pay", "vs", "pay method")

	flag.Parse()

	fmt.Printf("Pay method %s \n", *payMethod)

	strategyChosen := ChoiseStrategyToPay(*payMethod)
	strategyChosen.Pay()
}
