package main

import (
	"fmt"
)

// syntax using type declaration, followed by name and struct
type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	publicPlace string
	number      uint8
}

func main() {

	var u user
	u.name = "Aline"
	u.age = 21
	fmt.Println(u)

	// Another way to use struct
	// Is it possible to have struct inside struct
	addressExample := address{"Street", 25}
	userTwo := user{"Aline", 21, addressExample}
	fmt.Println(userTwo)

	// Another way to use struct without all the data
	// The field without information has a blank result
	userThree := user{name: "Aline"}
	fmt.Println(userThree)

}
