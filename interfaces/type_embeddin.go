package main

import "fmt"

type User struct {
	Name string
}

func (*User) notify() {
	fmt.Println("User notified")
}

type Admin struct {
	User
	Level int
}

// GO promueve los metodos no hace falta escribir admin.User.notify()

// si admin tiene su notidy,ocurre un shadowing y se llama al metodo de admin
func (a *Admin) notify() {
	fmt.Println("Admin notified")
}
func main() {
	admin := Admin{
		User: User{
			Name: "John",
		},
		Level: 1,
	}
	admin.notify()
}
