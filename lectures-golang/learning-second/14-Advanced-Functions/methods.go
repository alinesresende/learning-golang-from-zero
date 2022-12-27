package main

import "fmt"

// Methods
// Go supports methods defined in structs.

type user struct {
	name string
	age  uint
}

// The method contains a receiver argument
// Methods can be defined for any type of pointer or value receiver.

func (u user) save() {
	fmt.Println("save user", u.name)
}

func main() {
	user1 := user{"Aline", 26}
	user2 := user{"Alexandre", 28}

	user1.save()
	user2.save()

}
