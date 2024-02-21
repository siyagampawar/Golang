package main

import "fmt"

func main() {

	var num1 int
	var num2 int
	var num3 int

	fmt.Printf("Enter The Three Numbers : ")
	fmt.Scanf("%d %d %d", &num1, &num2, &num3)

	if (num1 > num2) && (num1 > num3) {
		fmt.Printf("\n%d is Greatest", num1)
	}
	if (num2 > num1) && (num2 > num3) {
		fmt.Printf("\n%d is Greatest", num2)
	} else {
		fmt.Printf("\n%d is Greatest", num3)
	}
}
