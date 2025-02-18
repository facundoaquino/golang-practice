package main

import "fmt"

type UnStruct struct {
	Nombre string
}

var unStruct UnStruct = UnStruct{Nombre: "Un nombre"} // Variable global

func CambiaNombre(u *UnStruct) {
	u.Nombre = "Otro nombre"
}

func main() {

	CambiaNombre(&unStruct)

	fmt.Println(unStruct.Nombre) // Imprime: Un nombre

}
