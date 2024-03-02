package main

import "fmt"

func sum(number int) int {

	if number <= 0 {
		return 0
	} else {
		return number + sum(number-1)
	}
}

func main() {
	var num int
	fmt.Println("Enter a number : ")
	fmt.Scanf("%d", &num)
	var result = sum(num)
	fmt.Println("sum is : ", result)
}
