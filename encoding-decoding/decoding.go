package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

type Animal struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	filename, _ := filepath.Abs("encoding-decoding/animals.json")

	data, err := os.ReadFile(filename)
	if err != nil {
		slog.Error("Error reading file... ", err)
	}

	var animals []Animal
	json.Unmarshal(data, &animals)
	fmt.Printf("animales decodificados desde json :  %+v", animals)

}
