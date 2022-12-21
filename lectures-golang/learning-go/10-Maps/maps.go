package main

import "fmt"

func main() {
	// inside the "[]" type of the curly braces
	// outside the "[]" type of the values
	// note: both fields need to be the same type
	// Example: map[string]string
	// it's forbidden: map[string]int
	user := map[string]string{
		"name":     "Pedro",
		"lastname": "Silva",
	}
	fmt.Println(user)

	// To access
	fmt.Println(user["name"])

	// Nested map
	user2 := map[string]map[string]string{
		"name": {
			"first": "Aline",
			"last":  "Resende",
		},
		"course": {
			"name":   "Engenharia",
			"fildes": "fildes1",
		},
	}
	fmt.Println(user2)

	// To delete
	delete(user2, "name")
	fmt.Println(user2)
}
