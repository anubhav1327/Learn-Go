package main

import "fmt"

const (
	first = iota
	second
	third = 'a'
	fourth
	fifth
)

func main() {
	fmt.Println("Conditional Statement is here")

	if third == 97 {
		fmt.Println("true 1")
	} else {
		fmt.Println("false 1")
	}

	if fourth == 97 {
		fmt.Println("true 2")
	}

	if first == 97 {
		fmt.Println("true 3")
	} else if second == 97 {
		fmt.Println("true 4")
	} else {
		fmt.Println("false 2")
	}
}
