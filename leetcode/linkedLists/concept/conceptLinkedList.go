package main

import "fmt"

/*
	linked list is a linear data structure where each element called node
	contains value, and  pointer(to the next node)


	there are multiple types:
	1. singly
	2. doubly
	circular linked list
*/

// singly
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node3 := &ListNode{Val: 3, Next: nil}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}
	head := &ListNode{Next: node1}

	for curr := head; curr != nil; curr = curr.Next {
		fmt.Println(curr.Val)
	}

	// head -> node1
}
