package main

import "fmt"

func main() {
	var imAnInt int
	var imAString string
	var imFloating float64
	var imEitherTrueOrFalse bool

	fmt.Printf("%T\n", imAnInt)
	fmt.Printf("%T\n", imAString)
	fmt.Printf("%T\n", imFloating)
	fmt.Printf("%T\n", imEitherTrueOrFalse)
}
