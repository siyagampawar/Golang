package main

import "fmt"

func addNumbers(n1 int, n2 int) (int, int) {

	sum := n1 + n2
	difference := n1 - n2
	return sum, difference
}

func main() {

	var n1 int
	var n2 int
	fmt.Println("Enter two numbers : ")
	fmt.Scanf("%d %d", &n1, &n2)
	sum, difference := addNumbers(n1, n2)
	fmt.Println("Sum : ", sum, "Difference : ", difference)
}
