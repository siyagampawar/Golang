package main

import "fmt"

func main() {

	squaredNumber := map[int]int{2: 4, 3: 9, 4: 16}

	for number, squared := range squaredNumber {
		fmt.Printf("Square of %d is %d\n", number, squared)
	}
}
