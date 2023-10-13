package main

import "fmt"

func main() {
	fmt.Println("Function by reference")

	a := a()
	b := b()
	sum(a, b)
	passbyref(&a, &b)
	variadic(1, 2, 3, 4, 5, 56)
	// initfun()
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

// Functions type

// 1- Pass by value
// 2- Pass by reference
// 3- Anonymous Function
// 4- Variadic function
// 5- init Function

// variadic function -chahe jitni value di jaye sab accept kar lega or sabse pahle run hoga mtlb jab ye run ho jayega tabhi baki code run hoga to ham ispe api fetch krke jab api aa jaye ya jab sab ready ho jaye tab baki ka code run kr dakte hai ise use krke

func variadic(v ...int) {
	fmt.Println("element are", v)
}

// init function

func init() {
	fmt.Println("This is init function - na koi value input lega na hi koi value return dega")
}
