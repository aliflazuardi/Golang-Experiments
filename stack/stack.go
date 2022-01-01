package main

import "fmt"

type Stack []string

// Check if the stack is empty
func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

// Pushing element to stack
func (s *Stack) Push(element string) {
	*s = append(*s, element)
}

// Pop (delete and return) top element of stack. return false if empty
func (s *Stack) Pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top element
		element := (*s)[index] // Get the element value
		(*s) = (*s)[:index]    // remove the top element by slicing it off
		return element, true
	}
}

func main() {
	var stack Stack

	stack.Push("Hello")
	fmt.Printf("Stack: %v\n", stack)
	stack.Push("Hi")
	fmt.Printf("Stack: %v\n", stack)
	stack.Push("Hey")
	fmt.Printf("Stack: %v\n", stack)

	for len(stack) > 0 {
		x, y := stack.Pop()
		if y {
			fmt.Printf("Element popped: %s\n", x)
		}
	}

}
