package main

import (
	"fmt"
	"unicode"
)

func main() {

	var mapCharacterCount map[string]int = make(map[string]int)

	var text string = "Hello, World!"

	for _, character := range text {

		if unicode.IsLetter(character) {

			mapCharacterCount[string(character)] = mapCharacterCount[string(character)] + 1
		}
	}

	fmt.Println(mapCharacterCount)

}
