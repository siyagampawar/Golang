package main

import "fmt"

func main() {
	time := 0
	if time > 0 {
		if time <= 10 {
			fmt.Println("Good Morning")
		} else if time <= 20 {
			fmt.Println("Good Day")
		} else {
			fmt.Println("Good Evening")
		}
	} else {
		fmt.Println("Invalid Time format")
	}
}
