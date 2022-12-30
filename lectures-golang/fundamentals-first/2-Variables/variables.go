package main

import "fmt"

func main() {
	//declare variable explicitly
	var variavel string = "Variavel1"
	fmt.Println(variavel)

	//declare variable implicitly
	variavel1 := "Variavel2"
	fmt.Println(variavel1)

	//Another way to declare a variable
	var (
		variavel3 string = "variavel3"
		variavel4 string = "variavel4"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel5", "variavel6"
	fmt.Println(variavel5, variavel6)

	//applies the same thing to constants
}
