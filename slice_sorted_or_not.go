package main

import (
	"fmt"
	"sort"
)

func main() {

	var status bool = false

	slice1 := []int{70, 20, 80, 40, 50, 60, 10, 30, 90, 100}
	slice2 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	status = sort.IntsAreSorted(slice1)

	if status == true {
		fmt.Println("Slice 1 is sorted")
	} else {
		fmt.Println("Slice 1 is not sorted")
	}
	status = sort.IntsAreSorted(slice2)

	if status == true {
		fmt.Println("Slice2 is sorted")
	} else {
		fmt.Println("Slice2 is not sorted")
	}

}
