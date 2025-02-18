package main

import (
	"fmt"
)

func fibonacci(n int) (valuePosition int) {

	fibonacciStart := []int{0, 1}

	if n == 0 {
		valuePosition = 0
	}

	if n == 1 {
		valuePosition = 1
	}

	for i := 2; i <= n; i++ {
		fibonacciAuxiliar := fibonacciStart[0] + fibonacciStart[1]
		fibonacciStart[0] = fibonacciStart[1]
		fibonacciStart[1] = fibonacciAuxiliar

	}
	valuePosition = fibonacciStart[1]

	return
}

func main() {
	fmt.Println(fibonacci(1100002212))
}
