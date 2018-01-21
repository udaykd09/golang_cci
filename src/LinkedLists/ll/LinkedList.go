package LinkedLists

import (
	"fmt"
)

type Node struct {
	Next *Node
	Prev *Node
	Key interface{}
}

type List struct {
	Head *Node
	Tail *Node
	Size int
}

func (L *List) InsertFront(key interface{}) {
	new := &Node{
		Next: L.Head,
		Prev: nil,
		Key: key,
	}
	if L.Head != nil {
		L.Head.Prev = new		
	}
	L.Head = new
	L.Size++
}

func (L *List) Display() {
	l := L.Head
	for l != nil {
		fmt.Printf("%v ->", l.Key)
		l = l.Next
	}
	fmt.Println()
}

func (L *List) Delete(key interface{}) {
	node := L.Head
	for node != nil {
		if node.Next.Key == key {
			node.Next = node.Next.Next
			break
		}
	L.Size--
	}
}
	