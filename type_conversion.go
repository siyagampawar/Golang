package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var avg float64

	num1 = 10
	num2 = 5
	avg = (float64(num1) + float64(num2)) / 2

	fmt.Printf("Average is %0.2f", avg)
}
