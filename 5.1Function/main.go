package main

import "fmt"

func main() {
	fmt.Println("Function by reference")

	a := a()
	b := b()
	sum(a, b)
	passbyref(&a, &b)
}

func a() int {
	return 10
}
func b() int {
	return 11
}

func sum(a, b int) {
	fmt.Println("Pass by value = ", a+b)
}

func passbyref(a, b *int) {
	fmt.Println("Pass by reference = ", *a+*b)
}
