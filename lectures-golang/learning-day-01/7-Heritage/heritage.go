package main

import "fmt"

type people struct {
	name     string
	lastName string
	age      uint8
	height   uint8
}

// writes struct name without passing types
type student struct {
	people  //
	course  string
	college string
}

func main() {
	p1 := people{"João", "Pedro", 20, 178}
	fmt.Println(p1)
	s1 := student{p1, "Engenharia", "USP"}
	fmt.Println(s1)

	// access form or struct
	fmt.Println(s1.name)
	fmt.Println(s1.height)
}
