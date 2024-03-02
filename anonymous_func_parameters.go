package main

import "fmt"

func main() {

	var sum = func(n1 int, n2 int) {
		result := n1 + n2
		fmt.Println("SUM IS : ", result)
	}

	sum(5, 3)
	fmt.Printf("%T", sum)
}
