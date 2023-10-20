package main

import (
	"fmt"
	"time"
)

/*

Goroutine :-

1:- every process which can concurrently in golang is called goroutine
2:- light weighted thread
3:- creation cost of goroutine is very small as compare to thread
4:- every program had at least single goroutine called main function
5:- when main is terminated then all goroutine will terminated means all routine works under main

syntax

func functionName(){
	// business logic
}

using go keyword as the prefix of your function call you can create goroutine
like = go functionName()


*/

func main() {
	fmt.Println("This is about Go rutine")
	go doSomeMagic()
	doSomeMagic()
	// time1()
	go func() {
		fmt.Println("go routine ")

	}()
	time.Sleep(1 * time.Second)
}

func doSomeMagic() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("This is ", i)
	}
}

// func time1() {
// 	for i := 1; i <= 61; i++ {

// 		if i > 60 {
// 			time1()
// 		} else {
// 			fmt.Println(i)
// 			time.Sleep(1 * time.Minute)
// 		}
// 	}
// }
