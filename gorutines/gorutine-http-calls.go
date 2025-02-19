package main

import (
	"encoding/json"
	"fmt"
	"mi_proyecto/utils"
	"os"
	"sync"
)

// con gorutinas para 50 llamadas http
// real    0m1.644s
// user    0m0.015s
// sys     0m0.046s

func main() {
	fmt.Println("Iniciando llamadas HTTP")
	var responses []map[string]interface{}
	newFile, _ := os.Create("responses.json")

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(50)

	for id := 1; id <= 50; id++ {
		fmt.Println("Procesando ID:", id)

		// a fines practicos ignoramos el error
		go getData(id, &wg, &responses, &mu)

	}

	wg.Wait()
	fmt.Println(responses)

	json.MarshalIndent(responses, "", "  ")

	defer newFile.Close()

	encoder := json.NewEncoder(newFile)
	encoder.SetIndent("", "  ")
	encoder.Encode(responses)

}

func getData(id int, wg *sync.WaitGroup, responses *[]map[string]interface{}, mu *sync.Mutex) {
	defer wg.Done()
	data, _ := utils.MakeAsyncHttpCall(id)
	mu.Lock()
	*responses = append(*responses, data)
	mu.Unlock()
}
