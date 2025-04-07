package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// append nodes to the list
func appendNode(head *ListNode, val int) *ListNode {
	newNode := &ListNode{Val: val}

	if head == nil {
		return newNode // first node
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode

	return head
}

func printlist(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d ->", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {

	var head *ListNode

	for _, val := range []int{1, 2, 3, 4} {
		head = appendNode(head, val)
	}

	fmt.Print("Original List: ")
	printlist(head)
}
