package main

import "fmt"

func main() {
	var name string
	fmt.Println("What's your name?")
	fmt.Print("Please enter here:")
	fmt.Scan(&name)
	fmt.Println()
	fmt.Println("Hello", name)
}
