package main

import "fmt"

func addNumbers(n1 int, n2 int) int {

	sum := n1 + n2
	return sum
}

func main() {
	sum := addNumbers(11, 22)
	fmt.Println("Sum : ", sum)
}
