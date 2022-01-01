package main

import "fmt"

func enqueue(queue []int, element int) []int {
	// Append the new element
	queue = append(queue, element)
	fmt.Println("Queued: ", element)

	return queue
}

func dequeue(queue []int) []int {
	// Extract the first element to dequeue
	element := queue[0]
	fmt.Println("Dequeued: ", element)

	return queue[1:] // Slice off the dequeued element
}

func main() {
	fmt.Println("Queue Tutorial Golang")

	// Implementing queue with slices
	var queue []int
	fmt.Printf("Queue: %v\n", queue)

	queue = enqueue(queue, 10)
	fmt.Printf("Queue: %v\n", queue)

	queue = enqueue(queue, 20)
	fmt.Printf("Queue: %v\n", queue)

	queue = dequeue(queue)
	fmt.Printf("Queue: %v\n", queue)

	queue = dequeue(queue)
	fmt.Printf("Queue: %v\n", queue)
}
