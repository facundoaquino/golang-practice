package utils

import (
	"math/rand"
	"time"
)

// Person struct
type Person struct {
	Name     string
	LastName string
	Age      int
}

// GenerateRandomPersons returns 50 random persons
func GenerateRandomPersons() []Person {
	firstNames := []string{"John", "Jane", "Alex", "Emily", "Chris", "Katie"}
	lastNames := []string{"Smith", "Johnson", "Williams", "Jones", "Brown", "Davis"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	persons := make([]Person, 50)
	for i := range persons {
		persons[i] = Person{
			Name:     firstNames[r.Intn(len(firstNames))],
			LastName: lastNames[r.Intn(len(lastNames))],
			Age:      r.Intn(71) + 10, // Age between 10 and 80
		}
	}
	return persons
}

// GENERATE A SLICE OF n [] INT
func GenerateRandomInts(n int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ints := make([]int, n)
	for i := range ints {
		ints[i] = r.Intn(100)
	}
	return ints
}
