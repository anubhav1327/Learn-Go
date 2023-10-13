package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("String Methods")

	var str1 string = "      @@@Anubhav Singh @@"
	fmt.Println(str1)
	tstr := strings.Trim(str1, "@")
	tstr2 := strings.TrimLeft(str1, "@")
	tstr3 := strings.TrimRight(str1, "@")
	tstr4 := strings.TrimSpace(str1)
	tstr5 := strings.TrimSuffix(str1, "@")
	tstr6 := strings.TrimPrefix(str1, "@")
	fmt.Println("total Trim", tstr)
	fmt.Println("TrimLeft", tstr2)
	fmt.Println("TrimRight", tstr3)
	fmt.Println("TrimSpace", tstr4)
	fmt.Println("TrimSuffix", tstr5)
	fmt.Println("TrimPrefix", tstr6)

}
