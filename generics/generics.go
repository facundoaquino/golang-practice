package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func returnSame[T any](some T) T {
	return some
}

func stringOrInt[T constraints.Integer | ~string](some T) {
	switch typeValue := any(some).(type) {
	case int:
		fmt.Println("This is an integer", typeValue)
	case string:
		fmt.Println("This is a string", typeValue)
	default:
		fmt.Println("This is a default for: ", typeValue)
	}
}

func main() {

	// puedo crear un returnSame para el tipo de dato especifico fsom ahora solo va a aceptar enteros
	fsom := returnSame[int]
	fmt.Println(fsom(1))

	fmt.Println(returnSame(1))
	fmt.Println(returnSame("Hello"))

	stringOrInt(1)
	stringOrInt(15.5)

}
