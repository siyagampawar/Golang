package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()

	switch today.Day() {
	case 5:
		fmt.Println("Clean the house")
	case 10:
		fmt.Println("buy grocery")
	case 15:
		fmt.Println("pay bills")
	case 20:
		fmt.Println("repair damaged stuff")
	case 25:
		fmt.Println("Go shopping")
	case 31:
		fmt.Println("Party")
	default:
		fmt.Println("Invalid Day")
	}
}
