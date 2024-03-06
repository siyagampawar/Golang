package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []string{"Honesty", "Is", "The", "Best", "Policy"}

	sort.Strings(slice)

	fmt.Println("Sorted Slice :")

	for _, item := range slice {
		fmt.Printf("%s ", item)
	}
}
