package main

import "fmt"

func main() {

	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	intSlice := arr[2:5]

	fmt.Println("Slice Elements: ")

	for index, ele := range intSlice {
		fmt.Printf("Index = %d,element=%d\n", index, ele)
	}
	for _, ele := range intSlice {
		fmt.Printf("element=%d\n", ele)
	}
	for index, _ := range intSlice {
		fmt.Printf("Index = %d\n", index)
	}
}
