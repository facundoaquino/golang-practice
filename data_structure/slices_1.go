package main

import "fmt"

func main() {

	sliceN := make([]int, 0)

	newSlice := append(sliceN, 1, 2, 34)

	fmt.Printf("Slice %v\n", newSlice)

	fmt.Printf("Slice %v", sliceN)

	// creating with index positions
	slicesStr := []string{100: "someStringTobegin"}
	fmt.Printf("Slice len %v\n", len(slicesStr))
	fmt.Printf("Slice  %v\n", slicesStr)

	// creating a nil slice of strings

	var nilSlice []string

	fmt.Printf("Slice len %v\n", len(nilSlice))

	//aumentando la cap con append

	simpleSlice := []string{"a", "b", "c", "d", "e"}

	newStringSlice := append(simpleSlice, "f", "g", "h")

	newStringSlice[0] = "z"

	fmt.Printf("Slice %v\n", simpleSlice)

	// es importante que si el slice tiene capacidad llena, con append, relocaliza la memoria y crea un nuevo array subyacente
	simpleSlice2 := make([]string, 0, 2)

	simpleSlice2 = append(simpleSlice2, "f")
	simpleSlice22 := append(simpleSlice2, "g")
	simpleSlice22[0] = "s"
	fmt.Printf("Slice %v\n", simpleSlice2)

	fmt.Printf("La capacidad de simpleSlice2 es %v\n", cap(simpleSlice2))
	fmt.Printf("Direccion de mem de slic2: %p  y de slice22 %p", &simpleSlice2[0], &simpleSlice22[0])

	//***********
	sliceToFunc := make([]int, 10, 10)
	fmt.Printf("slice mem %p\n", &sliceToFunc[0])
	sliceinFunc(sliceToFunc)
}

func sliceinFunc(slice []int) {
	fmt.Printf("slice mem in func %p\n", &slice[0])
}
