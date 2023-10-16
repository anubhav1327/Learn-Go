package main

import (
	"fmt"
	"reflect"
)

// Array are similar like other language go, java and c++
// Array containing similar type of data
// It has fixed length
// Due to fixed length they are not much popular in Golang
// Fixed length problem solved by 'Slice'
// Want to access the array element the use [] index operator
// Array starts from 0 to end array(len -1)

func main() {
	fmt.Println("This is all about array")

	//Types of array :- 1-one dimenssion    ,   2- multi dimenssion

	// we can assign array :- by var keyword and another method is sorthand method ( arr := array)

	//  type 1

	var arr1 [2]int /// for finding length = len(arr1)

	// type 2

	arr2 := [5]string{}

	arr1[0] = 23
	arr1[1] = 24

	arr2[0] = "Anubhav"
	arr2[1] = "singh"
	arr2[2] = "Mahi"
	arr2[3] = "Singh"
	arr2[4] = "na"

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println()

	fmt.Println(arr2, arr1)

	// By loops we can access every element of array like other language

	for i := 0; i < 2; i++ {
		fmt.Println(arr1[i], reflect.TypeOf(arr1))
	}

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i], reflect.TypeOf(arr2))
	}

	// Array by using slice method

	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9} /// Q- ye array bola jayega ya slice  || ans = ye slice nhi hai ye ek array hi hai bus isme fixed side utna hi hoga jitne array me element hoge

	fmt.Println(arr3)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i], reflect.TypeOf(arr3))
	}

	/// Golang me ham array me comperision kr sakte hai

	a1 := [3]int{1, 2, 3}
	a2 := [...]int{1, 2, 3}
	a3 := [3]int{4, 2, 3}

	fmt.Println(a1 == a2)
	fmt.Println(a2 == a3)
	fmt.Println(a3 == a1)

}
