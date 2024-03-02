package main

import (
	"fmt"
	"time"
)

func main() {

	switch today := time.Now(); { //switch initailzer
	case today.Day() < 5:
		fmt.Println("Clean the house")
	case today.Day() <= 10:
		fmt.Println("buy grocery")
	case today.Day() > 15:
		fmt.Println("pay bills")
	case today.Day() == 25:
		fmt.Println("Go shopping")
	default:
		fmt.Println("Invalid Day")
	}
}
