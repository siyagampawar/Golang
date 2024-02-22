package main

import "fmt"

func square(num int) {

	square := num * num
	fmt.Printf("Square of %d is %d\n", num, square)
}

func main() {

	square(3)
	square(66)
	square(10)
}
