package main

import (
	"fmt"
)

// give a singly list reverse the list and given its head

// this is a classic problem that involves modifying the pointer to reverse
// the direction of the list
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var revserse *ListNode

	curr := head

	for curr != nil {
		nextTemp := curr.Next
		curr.Next = revserse
		revserse = curr
		curr = nextTemp
	}

	return revserse
}

func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println("nil")
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	fmt.Print("original List: ")
	printList(head)

	// Reverse the linked list
	reverseHead := reverseList(head)

	fmt.Print("Reverse List: ")
	printList(reverseHead)
}
