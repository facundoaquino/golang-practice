package main

import "fmt"

func main() {

	var sliceN []int

	sliceN = append(sliceN, 1, 2, 3, 4, 5)
	tryToChangeSlice(&sliceN)
	// changeSlice(sliceN)
	// takeAPorcion(sliceN)

	fmt.Printf("Slice: %v\n", sliceN)

}
func tryToChangeSlice(sliceN *[]int) {
	slice2 := make([]int, 0)
	slice2 = append(slice2, 1, 2, 3, 4, 5, 9, 34, 2)
	*sliceN = slice2
}

// los slices pasan por referencia
func changeSlice(sl []int) {
	// sl[0] = 100
	sl = append(sl, 1)
}

// UNA PORCION DE UN SLICE ES UNA REFERENCIA A UNA PORCION DE MEMORIA (puntero)
func takeAPorcion(sl []int) {
	sl = sl[0:3]

	sl[0] = 100
}
