package main 

import (
	"LinkedLists/ll"
)
/*
Remove Dups: Write code to remove duplicates from an unsorted linked list.
FOLLOW UP
How would you solve this problem if a temporary buffer is not allowed?
*/

func removeDups(L *LinkedLists.List) {
	m := make(map[interface{}]int)
	var tmp *LinkedLists.Node 
	n := L.Head
	for n != nil {
		if m[n.Key] == 1 {
			tmp.Next = n.Next
		} else {
			m[n.Key]++
			tmp = n
		}
		n = n.Next
	}
}

func removeDupsNoBuff(l *LinkedLists.List) {
	current := l.Head
	for current != nil {
		runner := current
		for runner.Next != nil {
			if runner.Next.Key == current.Key {
				runner.Next = runner.Next.Next
			} else {
				runner = runner.Next
			}
		}
		current = current.Next
	}
}

func main() {
	List := &LinkedLists.List{}
	List.InsertFront("A")
	List.InsertFront("A")
	List.InsertFront("A")
	List.InsertFront("D")
	List.Display()
	removeDupsNoBuff(List)
	List.Display()
}

