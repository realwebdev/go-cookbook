package main

import "fmt"

// Queue struct // FIFO
type Queue struct {
	element []int
}

// Enque adds an element to the end of the queue
func (q *Queue) Enque(value int) {
	q.element = append(q.element, value)
}

// Deque removes the first element from the queue or the front of the queue
func (q *Queue) Deque() (int, bool) {
	if len(q.element) == 0 {
		return 0, false // queue is empty
	}
	value := q.element[0]
	q.element = q.element[1:]
	return value, true
}

// Front returns the first element from the queue
func (q *Queue) Front() (int, bool) {
	if len(q.element) == 0 {
		return 0, false // queue is empty
	}

	return q.element[0], true
}

// Isempty checks if the queue is empty
func (q *Queue) Isempty() bool {
	return len(q.element) == 0
}

func main() {
	queue := &Queue{}

	// Enque elements
	queue.Enque(1)
	queue.Enque(2)
	queue.Enque(3)

	// print the queue
	fmt.Println(queue) // Output: &{[1 2 3]}

	// Deque elements
	fmt.Println(queue.Deque()) // Output: 1, true
	fmt.Println(queue.Deque()) // Output: 2, true

	//check front
	fmt.Println(queue.Front()) // Output: 0, false

	fmt.Println(queue.Deque()) // Output: 3, true

	//check if queue is empty
	fmt.Println(queue.Isempty()) // Output: true

	fmt.Println(queue.Deque()) // Output: 0, false

}
