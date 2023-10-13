package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func main() {
// 	newinput := bufio.NewReader(os.Stdin)

// 	fmt.Println("Enter About You -")

// 	final_input, _ := newinput.ReadString('\n')

// 	fmt.Println("About Myself ->", final_input)
// }

func main() {
	fmt.Println("Welcome God")
	fmt.Println("Please god bless me")

	input := bufio.NewReader(os.Stdin)

	final_input, _ := input.ReadString('\n')

	fmt.Println("Beta You have to work hard ", final_input, "more time")

	inputBless, err := strconv.ParseFloat(strings.TrimSpace(final_input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to the Bless", inputBless+1)
	}
}
