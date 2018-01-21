package main 

import (
	"LinkedLists/ll"
)

// R	everse the linked list

func reverse(l *LinkedLists.List) {
	current := l.Head
	var prev *LinkedLists.Node
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}

func main() {
	List := &LinkedLists.List{}
	List.InsertFront("A")
	List.InsertFront("B")
	List.InsertFront("C")
	List.InsertFront("D")
	List.Display()
	reverse(List)
	List.Display()
}

