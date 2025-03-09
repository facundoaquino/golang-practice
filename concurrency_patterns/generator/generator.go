package main

import (
	"fmt"
	"time"
)

// GENERADORES CON CANALES

func main() {
	ch := boringFunc("boring")

	for i := 0; i < 3; i++ {

		fmt.Println(<-ch)
	}
}

func boringFunc(msg string) <-chan string {

	channelStr := make(chan string)
	go func() {

		for {
			channelStr <- fmt.Sprintf("%s", msg)
			time.Sleep(time.Second * 2)
		}
	}()
	return channelStr

}
