package main

import "fmt"

func countDown(number int) {
	if number > 0 {
		fmt.Println(number)
		countDown(number - 1)

		//PRINTS REVERSE
		// countDown(number - 1)
		// fmt.Println(number)

	} else {
		fmt.Println("CountDown Stops")
	}

}

func main() {
	fmt.Println("Countdown starts : ")
	countDown(10)
}
