package main

import "fmt"

// lainterface puede definir conjuntos de tipos tambien, solo se puede usar en restrinciones de tipo
type typeSetInterface interface {
	int | string | bool
}

func printTypeValue[T typeSetInterface](value T) {
	fmt.Println(value)
}

func main() {

	value := 2
	printTypeValue(value)
}
