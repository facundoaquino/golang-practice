package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type List struct {
	Head *Node
}

func (l *List) AddToTheEnd(n int) {

	newNode := &Node{Value: n}

	aux := l.Head

	if aux == nil {
		l.Head = newNode
		return
	}

	for aux.Next != nil {
		aux = aux.Next
	}

	aux.Next = newNode
}

func (l *List) AddToBeginig(n int) {
	newNode := &Node{Value: n}
	newNode.Next = l.Head
	l.Head = newNode
}

func (l List) PrintElements() {

	aux := l.Head

	for aux != nil {

		fmt.Printf("%v -> ", aux.Value)

		aux = aux.Next
	}

}

func main() {

	var list List

	list.AddToBeginig(1)
	list.AddToBeginig(2)
	list.AddToBeginig(3)
	list.AddToBeginig(4)

	list.PrintElements()

}
