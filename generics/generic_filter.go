package main

import "fmt"

type Person struct {
	Name     string
	LastName string
}

type Slice[E any] []E

func (s Slice[E]) Filter(keep func(E) bool) Slice[E] {

	var returned Slice[E]

	for _, value := range s {
		if keep(value) {
			returned = append(returned, value)

		}
	}
	return returned
}

func main() {
	var personSlice Slice[Person]

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

	var filtered []Person

	personSlice = perons
	filtered = personSlice.Filter(func(p Person) bool {
		return p.LastName == "perez"
	})

	fmt.Printf("Persons filtered: %+v", filtered)

}
