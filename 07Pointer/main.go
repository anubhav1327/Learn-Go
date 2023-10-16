package main

import "fmt"

func main() {
	fmt.Println("Pointers for GOlang")
	var a *int
	var b *int
	a = ptr(10)
	b = ptr(200)
	c := *a + *b        // yaha agar *c leta hu to error kyu aata hai
	fmt.Println(ptr(c)) // or agar upar *c lu ot niche *c return kru to error kyu aata hai
}

func ptr(v int) *int { // yaha *v kyu nhi le sakte
	return &v // agar upar *v lete hai to niche *v return kyu nhi kar sakte ham
}
