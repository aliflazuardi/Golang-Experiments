package main

import "fmt"

func main() {
	fmt.Println("Hashmap Tutorial Golang")

	// Create a hashmap make(map[key-type]val-type)
	m := make(map[string]string)

	// initialize key-value
	m["foo"] = "bar"
	m["hey"] = "hello"

	// Print hashmap and value from a key
	fmt.Printf("Map: %v\n", m)
	fmt.Printf("Map with foo key: %v\n", m["foo"])

	// Extracting values
	v1 := m["foo"]
	fmt.Printf("Value 1: %s\n", v1)

	// Get the length of the hashmap
	fmt.Printf("Map Length: %v\n", len(m))

	// Delete key-value pair from map
	delete(m, "hey")
	fmt.Printf("Map after delete: %v\n", m)

	// Check if key is present
	_, prs := m["hey"]
	fmt.Printf("is Present? %v\n", prs)

	// Delcare and initialize map in one line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)

	// Iterate through map
	for k, v := range n {
		fmt.Printf("%s: %v\n", k, v)
	}

}
