package main

import (
	u "mi_proyecto/algorithms_exercises"
)

func main() {

	newList := u.GeneratedListForNValues(3)

	newListR := reverseList(newList)
	u.PrintValuesFromList(newListR)
}

func reverseList(listHead *u.ListNode) *u.ListNode {
	var newHead *u.ListNode
	auxNode := listHead

	for auxNode != nil {
		next := auxNode.Next
		auxNode.Next = newHead
		newHead = auxNode
		auxNode = next
	}

	return newHead
}
