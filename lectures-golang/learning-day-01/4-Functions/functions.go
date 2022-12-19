package main

import "fmt"

// After the parameters specifies the return of the function
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)
}
