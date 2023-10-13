package main

import (
	"bufio"
	"fmt"
	"os"
)

// func main() {
// 	fmt.Println("Ganpati")
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter Your Number-")

// 	// comma ok  // err ok syntax

// 	input, _ := reader.ReadString('\n')

// 	fmt.Println("My number is -> ", input)
// }

func main() {
	newinput := bufio.NewReader(os.Stdin)

	fmt.Println("hello Go Learner")
	fmt.Println("Enter your full name->")

	final_input, _ := newinput.ReadString('\n')
	fmt.Printf("")

	fmt.Println("My full name is ", final_input)
	fmt.Println("Now you are varified for learning Golang")
}
