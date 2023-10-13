package main

import "fmt"

func main() {
	a := 22
	b := 44

	fmt.Println(test(a, b))
}

// function syntax = isme return hi krna hota hai isme function ke andar print command nhi use ka r sakte hai
// and synatax me phle bricket me variables define kre then next bricket me return type define kre 2 return use
// kr sakte hai ya ek bhi like ham return me likh sakte hai ki

// func function name(first variable, second variable)
func test(num1 int, num2 int) (int, error) {
	return num1 + num2, nil
}

/// second way

// func test(num1 int, num2 int) int {
// 	return num1 + num2
// }
