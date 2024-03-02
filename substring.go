package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Golang is similar to C"
	var substr string = "lang"

	if strings.Contains(str, substr) {
		fmt.Printf("String (%s) contains substring (%s)", str, substr)
	} else {
		fmt.Printf("String (%s) does not contain substring (%s)", str, substr)
	}
}
