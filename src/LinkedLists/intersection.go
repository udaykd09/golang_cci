package main 

import (
	"LinkedLists/ll"
	"fmt"
)

/*
Intersection: Given two (singly) linked lists, determine if the two lists intersect. Return the
intersecting node. Note that the intersection is defined based on reference, not value. That is, if the
kth node of the first linked list is the exact same node (by reference) as the jth node of the second
linked list, then they are intersecting.
*/
func isIntersecting(l1, l2 *LinkedLists.List) bool {
	fmt.Println(l1.Size, l2.Size)
	return true
}

func main() {
	List := &LinkedLists.List{}
	List.InsertFront("A")
	List.InsertFront("B")
	List.InsertFront("C")
	List.InsertFront("D")
	List.Display()
	
	List2 := &LinkedLists.List{}
	List2.InsertFront("E")
	List2.Head.Next = List.Head.Next
	List2.InsertFront("F")
	List2.Display()
	
	fmt.Println(isIntersecting(List, List2))
}

