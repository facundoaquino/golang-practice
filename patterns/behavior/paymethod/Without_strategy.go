package main

import (
	"flag"
	"fmt"
)

func main() {

	payMethod := flag.String("pay", "vs", "pay method")

	flag.Parse()

	fmt.Printf("Pay method %s \n", *payMethod)

	switch *payMethod {
	case "visa":
		payWithvisa()
	case "master":
		payWithMaster()

	default:
		fmt.Println("Paying with default or incorrect method to pay")
	}
}

func payWithvisa() {
	fmt.Print("Paying with visa method.....")
}

func payWithMaster() {
	fmt.Print("Paying with master method.....")
}
