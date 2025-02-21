package main

import (
	"fmt"
)

func main() {

	// los canales con bufer se bloquean cuando estan llenos
	mychannel := make(chan string, 1)

	mychannel <- "Bob"
	username := <-mychannel
	mychannel <- "Bob"
	username = <-mychannel

	fmt.Println(username)
}

// podemos usar los canales para proteger el acceso a los datos (lectura escritura)

func sendMsg(ch chan<- string) {
	ch <- "Hello"
}

func readMsg(ch <-chan string) {
	msg := <-ch
	fmt.Println(msg)
}
