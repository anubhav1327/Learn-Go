package main

import "fmt"

func main() {
	fmt.Println("Closure in Golang")
	val := company()
	fmt.Println(val())
	fmt.Println(val())
	fmt.Println(val())
	fmt.Println(val())
	fmt.Println(val())
	fmt.Println(val())
	v := company()
	fmt.Println(v())
}

func company() func() int {
	a := 0
	return func() int {
		a++
		return a
	}
}
