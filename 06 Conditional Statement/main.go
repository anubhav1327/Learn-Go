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

	var anu = 98
	fmt.Println("Conditional Statement is here")
	// fmt.Println(anu)
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(fourth)
	fmt.Println(fifth)

	if anu == 90 {
		fmt.Println("Anu is present")
	} else {
		fmt.Println("Anu is absent")
	}

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
