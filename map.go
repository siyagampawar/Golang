package main

import "fmt"

func main() {
	subjectMarks := map[string]float32{"Golang": 10, "C": 50, "Python": 100}
	subjectMarks["Golang"] = 40
	subjectMarks["R"] = 70 // alphabetical order mein appemd karta hain
	delete(subjectMarks, "Golang")
	fmt.Println(subjectMarks)
	fmt.Println(subjectMarks["R"])
}
