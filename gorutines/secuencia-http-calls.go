package main

import (
	"encoding/json"
	"fmt"
	"mi_proyecto/utils"
	"os"
)

// 50 llamadas http secuenciales escribiendo las respuestas en un archivo JSON

// real    0m4.082s
// user    0m0.045s
// sys     0m0.030s

func main() {
	fmt.Println("Iniciando llamadas HTTP")
	var responses []map[string]interface{}
	newFile, err := os.Create("responses.json")

	if err != nil {
		fmt.Println("Error al crear archivo:", err)
		return
	}

	for id := 1; id <= 50; id++ {
		fmt.Println("Procesando ID:", id)

		// Hacer la llamada HTTP
		data, err := utils.MakeAsyncHttpCall(id)
		if err != nil {
			fmt.Println("Error en la solicitud:", err)
			continue // Si hay error, sigue con el siguiente ID
		}

		// Agregar la respuesta al slice
		responses = append(responses, data)
	}

	// Convertir el resultado a JSON y mostrarlo
	jsonOutput, err := json.MarshalIndent(responses, "", "  ")
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	defer newFile.Close()

	encoder := json.NewEncoder(newFile)
	encoder.SetIndent("", "  ") // Para que el JSON tenga formato legible
	encoder.Encode(responses)
	fmt.Println(string(jsonOutput))

}
