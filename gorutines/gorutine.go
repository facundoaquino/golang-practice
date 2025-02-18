package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

const finalToWrite = 10

func main() {
	var wg sync.WaitGroup
	wg.Add(finalToWrite) // Indicamos cuántas goroutines vamos a lanzar

	for i := 0; i < finalToWrite; i++ {
		go get(i, &wg) // Pasamos el WaitGroup a la goroutine
	}

	wg.Wait() // Esperamos a que todas las goroutines terminen
}

func get(userID int, wg *sync.WaitGroup) {
	defer wg.Done() // Indica que esta goroutine ha terminado

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(userID))
	if err != nil {
		fmt.Println("Error en la solicitud:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)

	body, _ := io.ReadAll(resp.Body)

	fmt.Println("Respuesta JSON:", string(body))
}
