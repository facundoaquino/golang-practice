package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

var usersCache = make(map[int]User)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	mux.HandleFunc("POST /users", createUser)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, World!")
}

func createUser(w http.ResponseWriter, r *http.Request) {

	var user User

	fmt.Println(r.Header)

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	usersCache[len(usersCache)+1] = user

	w.WriteHeader(http.StatusNoContent)

	fmt.Println(user)

	fmt.Fprintf(w, "User created")
}
