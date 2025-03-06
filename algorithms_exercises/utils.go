package algorithmsexercises

import "fmt"

type ListNode struct {
	Next  *ListNode
	Value int
}

// Generates a list for integer values ordener 1 ---> 2 ---> 3 ----> xxx
func GeneratedListForNValues(n int) *ListNode {

	var newList *ListNode

	for i := 1; i <= n; i++ {
		newList = AddValuesToListAtTheEnd(i, newList)
	}

	return newList
}

func AddValuesToListAtTheEnd(value int, headNode *ListNode) *ListNode {

	newNode := &ListNode{Value: value}

	if headNode == nil {
		return newNode
	}

	nodeAux := headNode
	for nodeAux.Next != nil {
		nodeAux = nodeAux.Next
	}
	nodeAux.Next = newNode

	return headNode

}

func PrintValuesFromList(list *ListNode) {

	for list != nil {
		fmt.Printf("Value: %v \n", list.Value)
		list = list.Next
	}
}
