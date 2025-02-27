package main

import "fmt"

type myFuncInt func(i int) int

// Los tipos genéricos en Go solo pueden usarse en funciones, structs e interfaces, pero no en alias de tipos de función.
// type myFuncGeneric func[t T](t T) T

// Go permite alias de tipos genéricos cuando el alias se basa en tipos de datos estructurados (como []E, map[K]V, struct{}), pero no en alias de funciones.
type SliceInt [][]int
type Slice[E any] []E

func main() {
	myFInt := func(i int) int {
		return i + 2
	}

	fmt.Printf("result from operations with ints: %d", myFInt(3))

}
