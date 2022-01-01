package main

import "fmt"

func main() {
	fmt.Println("Hashmap Tutorial Golang")

	// Create a hashmap make(map[key-type]val-type)
	m := make(map[string]string)

	// initialize key-value
	m["alif"] = "mifta"
	m["erlang"] = "dinda"

	// Print hashmap and value from a key
	fmt.Printf("Map: %v\n", m)
	fmt.Printf("Map with alif key: %v\n", m["alif"])

	// Extracting values
	v1 := m["alif"]
	fmt.Printf("Value 1: %s\n", v1)

	// Get the length of the hashmap
	fmt.Printf("Map Length: %v\n", len(m))

	// Delete key-value pair from map
	delete(m, "erlang")
	fmt.Printf("Map after delete: %v", m)

}
