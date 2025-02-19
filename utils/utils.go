package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func MakeAsyncHttpCall(id int) (map[string]interface{}, error) {
	// Hacer llamada HTTP
	var url string = "https://jsonplaceholder.typicode.com/todos/"

	resp, err := http.Get(url + strconv.Itoa(id))

	if err != nil {
		panic(err)
	}

	// el garbage collector puede tardar en desocupar los recursos que se le asiganron a la respuesta
	defer resp.Body.Close()

	var result map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	fmt.Println(result)

	return result, nil

}
