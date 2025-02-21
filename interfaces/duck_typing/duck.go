package main

import "fmt"

//  si dice cuak puede ser (es) un pato, en go no necesitamos implementar una interfaz explicitamente
//  para que un tipo sea considerado un pato, simplemente necesitamos que tenga los metodos necesarios

type Duck interface {
	QuackQuack()
}

type BirdDuckAnimal struct {
	Name  string
	Age   int
	Color string
}

func (b BirdDuckAnimal) QuackQuack() {
	fmt.Println("Cuak Cuak")
}

func main() {

	duck := BirdDuckAnimal{
		Name:  "Pato",
		Age:   2,
		Color: "Amarillo",
	}

	duck.QuackQuack()

}
