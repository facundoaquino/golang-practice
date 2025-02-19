package main

import (
	"fmt"
	"mi_proyecto/utils"
	"runtime"
	"sync"
)

// creacion de 2 slices random de 100000000 de elementos y operacion - ITERATIVO
// myRandomNumbers := utils.GenerateRandomInts(100000000)
// myRandomNumbers2 := utils.GenerateRandomInts(50000000)
// real    0m1.657s
// user    0m0.030s
// sys     0m0.015s

// func duplicateNumber2(numbers []int) {
// 	for index, number := range numbers {
// 		numbers[index] = number * 2
// 	}
// }

// func myGorutine(number *int, ch chan int, iteration int) {
// 	if iteration == 1 {
// 		ch <- 1
// 	}
// 	for i := 0; i < iteration; i++ {
// 		*number++
// 	}
// 	ch <- *number
// }

func duplicateNumber(numbers []int, wg *sync.WaitGroup) {
	for index, number := range numbers {
		numbers[index] = number * 2
	}
	wg.Done()

}

func generatedConcurretly(n int, ch chan []int, wg *sync.WaitGroup) {
	ch <- utils.GenerateRandomInts(n)
	wg.Done()
}

func main() {
	fmt.Println("Numero de CPUs: ", runtime.NumCPU())
	var sizeElements int = 200000000

	runtime.GOMAXPROCS(runtime.NumCPU())

	ch := make(chan []int)
	ch2 := make(chan []int)

	var wg sync.WaitGroup
	wg.Add(2)

	var myRandomNumbers []int
	var myRandomNumbers2 []int

	go generatedConcurretly(sizeElements, ch, &wg)
	go generatedConcurretly(sizeElements, ch2, &wg)

	myRandomNumbers = <-ch
	myRandomNumbers2 = <-ch2
	wg.Wait()

	fmt.Println(myRandomNumbers[1])
	wg.Add(2)

	go duplicateNumber(myRandomNumbers, &wg)
	go duplicateNumber(myRandomNumbers2, &wg)
	wg.Wait()

	fmt.Println(myRandomNumbers[1])

}
