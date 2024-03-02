package main

import "fmt"

func main() {
	var x interface{} //var x interface{} = "RKNEC" {Try}

	switch i := x.(type) {
	case nil:
		fmt.Printf("type of x:%T", i)
	case int:
		fmt.Printf("x is int")
	case float64:
		fmt.Printf("x is float64")
	case bool, string:
		fmt.Printf("x is bool or string")
	default:
		fmt.Printf("Don't know the type")
	}
}
