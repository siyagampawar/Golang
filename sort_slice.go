package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{70, 20, 80, 40, 50, 60, 10, 30, 90, 100}

	sort.Ints(arr)

	fmt.Println("Slice Elements: ")

	for _, ele := range arr {
		fmt.Printf("element=%d\n", ele)
	}

}
