package main

import "fmt"

// nivel de desreferencia doble: ya le esta pasando una copia del puntero
func changePointer(nPointer **int) {
	*nPointer = nil
}

func main() {

	var nPoint = new(int)

	*nPoint = 100

	fmt.Printf("Puntero: %v\n", nPoint)

	changePointer(&nPoint)

	fmt.Printf("Puntero: %v\n", nPoint)

	fmt.Println(*nPoint)

}
