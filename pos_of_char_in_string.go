package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "computer"
	var ind int

	ind = strings.Index(str, "t")
	if ind == -1 {
		fmt.Println("Character does not exist")
	} else {
		fmt.Println("Index is : ", ind)
	}

}
