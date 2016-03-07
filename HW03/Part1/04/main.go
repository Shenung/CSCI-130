package main

import "fmt"

func main() {
	var first, second int
	fmt.Println("Give me a number:")
	fmt.Scanln(&first)
	fmt.Println("Give me second number bigger than the first:")
	fmt.Scanln(&second)
	fmt.Println("Remainder = ", second%first)
}
