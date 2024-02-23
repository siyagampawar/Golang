package main

import "fmt"

var greet = func() {
	fmt.Printf("Hello,How are you")
}

func main() {

	// var greet = func() {
	// 	fmt.Printf("Hello,How are you")
	// }
	var welcome = greet
	greet()
	welcome()
}
