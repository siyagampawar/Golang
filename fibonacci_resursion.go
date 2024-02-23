package main

import "fmt"

func fibonacci(number int) int {

	if number == 1 || number == 2 {
		return 1
	} else {
		return fibonacci(number-1) + fibonacci(number-2)
	}
}
func main() {

	var n int
	fmt.Println("Enter a number : ")
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
}
