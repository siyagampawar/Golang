package main

import "fmt"

func main() {
	var x int = 10
	var y int = 20
	add(x, y)

}

func add(x int, y int) {
	z := x + y
	fmt.Println("Addition is : ", z)
}
