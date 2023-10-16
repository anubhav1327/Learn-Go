package main

import "fmt"

func main() {
	fmt.Println("Loops in GOlang")

	i := 23

	// for loop

	for j := 0; j < i; j++ {
		fmt.Println(j)
	}

	//infinite loop

	// for {
	// 	fmt.Println("Infinite lopp")
	// }

	// While Loop

	var k int = 2

	for k < 23 {
		fmt.Println(k)
		k++
	}

	// Range Loop

	variable := []string{"a", "b", "c", "d"}
	for l, value := range variable {
		fmt.Println(l, value)
	}
}

// git init
// git add README.md
// git commit -m "first commit"
// git branch -M main
// git remote add origin https://github.com/anubhav1327/dsds.git
// git push -u origin main
