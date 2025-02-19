package main

import (
	"fmt"
	"mi_proyecto/utils"
	"time"
)

func main() {

	start := time.Now()

	arr1 := utils.GenerateRandomInts(500000000)
	// fmt.Printf("Arr1: %v\n", arr1)
	arr2 := utils.GenerateRandomInts(500000000)
	arr3 := utils.GenerateRandomInts(500000000)
	arr4 := utils.GenerateRandomInts(500000000)
	arr5 := utils.GenerateRandomInts(500000000)

	sum := 0

	for i := 0; i < len(arr1); i++ {
		sum += arr1[i]
	}

	for i := 0; i < len(arr2); i++ {
		sum += arr2[i]
	}

	for i := 0; i < len(arr3); i++ {
		sum += arr3[i]
	}

	for i := 0; i < len(arr4); i++ {
		sum += arr4[i]
	}

	for i := 0; i < len(arr5); i++ {
		sum += arr5[i]
	}

	elapsed := time.Since(start)
	fmt.Println("Suma de arrs: ", sum)
	fmt.Println("Tiempo de ejecucion:  ", elapsed.Seconds(), " segundos")

}
