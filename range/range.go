package main

import (
	"fmt"
	"mi_proyecto/utils"
)

func main() {

	arrInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, value := range arrInt {
		println("Index: ", index, "Value: ", value)
	}

	var persons []utils.Person = utils.GenerateRandomPersons()
	fmt.Println("Persons: ", persons)

	// filtramos las personas mayores de edad con el filter

	olderPersons := filterPersons(persons, isMayorAge)
	fmt.Println("Older persons: ", olderPersons)

	changesAge(persons)

	fmt.Println("Persons: ", persons)
}

func isMayorAge(person utils.Person) bool {
	return person.Age >= 70
}

func changesAge(persons []utils.Person) {
	for i := range persons {
		persons[i].Age = 35
	}
}

// generic filter for persons

func filterPersons(persons []utils.Person, filter func(utils.Person) bool) (result []utils.Person) {

	for _, person := range persons {
		if filter(person) {
			result = append(result, person)
		}
	}

	return result
}
