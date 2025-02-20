package main

import "fmt"

func heapAlloc() *int {
	x := 42
	return &x // Posible escape al heap
}

func stackAlloc() int {
	y := 24
	return y // No escapa, se queda en el stack
}

func main() {
	a := heapAlloc()
	b := stackAlloc()
	fmt.Println(*a, b)
}
