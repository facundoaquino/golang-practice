package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {

	ctxParent := context.Background()

	ctxParent = context.WithValue(ctxParent, "someKey", "someValue")

	fmt.Printf("valor de la key en el context someKey: %s\n", ctxParent.Value("someKey"))

	var cancelAfterTo context.CancelFunc

	ctxParent, cancelAfterTo = context.WithTimeout(ctxParent, time.Second*2)

	defer cancelAfterTo()

	chanToReciveFinishProcess := make(chan struct{})

	go callToHardProccess(chanToReciveFinishProcess)

	select {
	case <-chanToReciveFinishProcess:

		log.Println("termino primero el proceso pesado")
	case <-ctxParent.Done():
		log.Println("Contexto termino por timeOut")
	}

}

func callToHardProccess(c chan<- struct{}) {
	time.Sleep(time.Second * 1)
	c <- struct{}{}
}
