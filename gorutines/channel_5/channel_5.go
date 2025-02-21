package main

import (
	"fmt"
)

type Server struct {
	users       map[string]string
	userChannel chan string
}

func NewServer() *Server {
	return &Server{
		users:       make(map[string]string),
		userChannel: make(chan string),
	}
}

func (s *Server) Start() {
	go s.Loop()
}

func (s *Server) Loop() {
	for {
		user := <-s.userChannel
		s.users[user] = user
		fmt.Printf("User %s added\n", user)
	}

}

func (s *Server) AddUser(username string) {
	s.users[username] = username
}

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
