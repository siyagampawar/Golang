package main

import "fmt"

func factorial(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * factorial(number-1)
	}
}
func main() {

	var n int
	fmt.Println("Enter a number : ")
	fmt.Scanf("%d", &n)
	var result = factorial(n)
	fmt.Println("Result is : ", result)
}
