package main

import "fmt"

func main() {

	var num1 int
	var num2 int
	var num3 int

	// fmt.Println("Enter The Three Numbers : ")
	// fmt.Scanf("%d %d %d", &num1, &num2, &num3)

	fmt.Printf("Enter The First Number : ")
	fmt.Scanf("%d", &num1)
	fmt.Printf("Enter The Second Number : ")
	fmt.Scanf("%d", &num2)
	fmt.Print("Enter The Third Number : ")
	fmt.Scanf("%d", &num3)

	// fmt.Println("Enter The Three Numbers 10: ")
	// fmt.Scanln("%d %d %d", &num1, &num2, &num3)

	if (num1 > num2) && (num1 > num3) {
		fmt.Printf("\n%d is Greatest", num1)
	}
	if (num2 > num1) && (num2 > num3) {
		fmt.Printf("\n%d is Greatest", num2)
	}
	if (num3 > num2) && (num3 > num1) {
		fmt.Printf("\n%d is Greatest", num3)
	}
}
