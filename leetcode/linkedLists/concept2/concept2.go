package main

import "fmt"

// create a linked list and add few ints

type LinkedList struct {
	Val  int
	Next *LinkedList
}

func appendLinkedList(head *LinkedList, val int) *LinkedList {
	newNode := &LinkedList{Val: val}

	if head == nil {
		return newNode // if the head is empty newnode becomes the head
	}

	// traverse the last node
	current := head
	for current.Next != nil {
		current = current.Next
	}

	// link the last node to the new node
	current.Next = newNode

	return head

}
func printLinkedList(head *LinkedList) {
	current := head
	for current != nil {
		fmt.Printf("%d ->", current.Val)
		current = current.Next
	}
	fmt.Println(nil)
}

func main() {
	var head *LinkedList
	a := []int{1, 2, 3, 4, 5, 6}

	for _, elements := range a {
		head = appendLinkedList(head, elements)
	}

	fmt.Print("list ")
	printLinkedList(head)
}
