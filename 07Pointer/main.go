package main

import "fmt"

func main() {
	fmt.Println("Pointers for GOlang")
	var a *int
	var b *int
	a = ptr(10)
	b = ptr(200)
	c := *a + *b
	fmt.Println(ptr(c))
}

func ptr(v int) *int {
	return &v
}
