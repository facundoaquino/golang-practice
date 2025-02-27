package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	response := make([]User, 0)

	var numberOfrequest = 5

	url := "https://jsonplaceholder.typicode.com/todos/"

	respChannel := make(chan User)

	for i := 1; i < numberOfrequest; i++ {
		go makeHttpCalle(i, url, respChannel)

	}

	for i := 1; i < numberOfrequest; i++ {

		response = append(response, <-respChannel)
	}
	fmt.Printf("responses: %+v \n", len(response))
}

func makeHttpCalle(i int, url string, respCh chan<- User) {
	//make http call
	fmt.Printf("Make %d call\n", i)
	resp, _ := http.Get(fmt.Sprintf("%s%d", url, i))
	//fmt.Printf("response: %v", resp)

	var userU User
	body, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &userU)
	fmt.Printf("User %d = %+v \n", i, userU)
	respCh <- userU
}
