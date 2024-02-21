package main

import "fmt"

func main() {
	// var a int = 30
	// var b int = 20
	// if a > b {
	// 	fmt.Printf("%d is greater than %d", a, b)
	// }
	// if 30 > 20 {
	// 	fmt.Printf("30 is greater than 20")
	// }
	// if !(a < b) {
	// 	fmt.Printf("will run always")
	// }

	x := 8

	// if y := 10; x < y {
	// 	fmt.Print("X is less than y")
	// }

	//fmt.Print(y)//gives error ,y was declared locally just for the if statement

	// if y := 10;a:=4; x < y {//gives error as we cannot give two statements for initialization
	// 	fmt.Print("X is less than y")
	// 	fmt.Print(y)
	// }

	if y, b := 10, 4; x < y { //doesnt give error as humne initialization ek hi line mein kar diya//if pehle statement ko as initialization leta hain but 2nd statement run karta hain this is why we cannot give two statements as initialization
		fmt.Print("X is less than y")
		fmt.Print(b)
	}
}
