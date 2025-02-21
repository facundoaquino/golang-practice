package main

import (
	"fmt"
	"math/rand"
)

/**
	SE VAN A EJECUTAR EN PARALELO LAS GORUTINES Y LA PRIMERA QUE GENERE UN NUMERO ALEATORIO 1, 2 O 3
	VA A SER LA QUE TERMINE PORQUE ESTA ESCUCHANDO UN SELECT

**/

func generateRandom(channel chan int) {
	fmt.Println("\nStarting gorutine")
	for {
		number := rand.Intn(500) + 1
		fmt.Printf("%d ", number)
		if number == 1 || number == 2 || number == 3 {
			fmt.Println("\nFinishing gorutine")
			channel <- number
			return
		}

	}
}

func main() {

	channel1 := make(chan int)
	channel2 := make(chan int)
	channel3 := make(chan int)

	go generateRandom(channel1)
	go generateRandom(channel2)
	go generateRandom(channel3)

	select {
	case number := <-channel1:
		if number == 1 {
			fmt.Println("\nChannel 1: 1")
		}
	case number := <-channel2:
		if number == 2 {
			fmt.Println("\nChannel 2: 2")
		}
	case number := <-channel3:
		if number == 3 {
			fmt.Println("\nChannel 3: 3")
		}
	}
}
