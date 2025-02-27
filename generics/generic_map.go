package main

import "fmt"

type Person struct {
	Name     string
	LastName string
}

func MyMapGenericFunction[K, V any](s []K, mapper func(K) V) []V {
	out := make([]V, len(s))

	for index, value := range s {
		out[index] = mapper(value)

	}

	return out
}
func main() {

	perons := []Person{
		{"maria", "gonzales"},
		{"maria", "gonzales"},
		{"federico", "grido"},
		{"julieta", "gonzales"},
		{"maria", "gonzales"},
		{"maria", "perez"},
		{"luciana", "ramirez"},
		{"maria", "vazquez"},
		{"maria", "gonzales"},
	}

	peronsMappered := MyMapGenericFunction(perons, func(per Person) string {

		return per.LastName

	})

	fmt.Printf("Personas mapeadas a apellido: %v", peronsMappered)
}
