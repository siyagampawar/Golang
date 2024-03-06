package main

import "fmt"

func main() {
	primeNum := []int{2, 3, 5, 7}
	numbers := []int{1, 2, 3}

	copy(numbers, primeNum)
	numbers = append(numbers, 5)
	fmt.Println("Numbers: ", numbers)
	fmt.Println(cap(numbers))
}
