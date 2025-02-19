package main

import "fmt"

func myGorutine(number *int, ch chan int, iteration int) {
	if iteration == 1 {
		ch <- 1
	}
	for i := 0; i < iteration; i++ {
		*number++
	}
	ch <- *number

}

func main() {

	ch := make(chan int)

	var number int

	// NO ESTOY MANEJANDO EL MANEJO CONCURRENTE A NUMBER QUE PUEDEN TENER AMBAS GORUTINES PERO ES PARA VER QUE TERMINA SIEMPRE PRIMERO LA DEL 1

	go myGorutine(&number, ch, 1)
	go myGorutine(&number, ch, 10000000)

	number = <-ch

	fmt.Println(number)
}
