package main

import "fmt"

func main() {

	var num int

	fmt.Println("Enter a Number : ")
	fmt.Scanf("%d", &num)

	if num%2 == 0 {
		fmt.Println("The number is EVEN")
	} else {
		fmt.Println("The number is ODD")
	}
}
