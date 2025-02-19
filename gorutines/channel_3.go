package main

import (
	"fmt"
	"mi_proyecto/utils"
	"time"
)

func sumArr(arr []int, ch chan int) {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	ch <- sum
}

func main() {

	start := time.Now()

	arr1 := utils.GenerateRandomInts(500000000)
	// fmt.Printf("Arr1: %v\n", arr1)
	arr2 := utils.GenerateRandomInts(500000000)
	arr3 := utils.GenerateRandomInts(500000000)
	arr4 := utils.GenerateRandomInts(500000000)
	arr5 := utils.GenerateRandomInts(500000000)

	var ch = make(chan int)

	go sumArr(arr1, ch)
	go sumArr(arr2, ch)
	go sumArr(arr3, ch)
	go sumArr(arr4, ch)
	go sumArr(arr5, ch)

	total := 0
	for i := 0; i < 4; i++ {
		total += <-ch
	}

	fmt.Println("Suma de arrs: ", <-ch)

	elapsed := time.Since(start)
	fmt.Println("Tiempo de ejecucion:  ", elapsed.Seconds(), " segundos")

}
