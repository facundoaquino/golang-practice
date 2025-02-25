package main

import "fmt"

func main() {
	var word string = "menem"

	var isPal bool = isPalindrome(word)

	fmt.Printf("Is palindrome?:   %t", isPal)
}

func isPalindrome(word string) bool {

	if len(word) == 1 || len(word) == 0 {

		return true
	}
	if word[0] == word[len(word)-1] {

		return isPalindrome(word[1 : len(word)-1])
	} else {
		return false
	}
}
