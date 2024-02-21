package main

import (
	"fmt"
)

func main() {

	var num float64
	var abs float64

	fmt.Println("Enter a Number : ")
	fmt.Scanf("%f", &num)

	if num < 0 {
		abs = -1 * num
	} else {
		abs = num
	}

	fmt.Printf("The Absolute Value is : %f", abs)
}
