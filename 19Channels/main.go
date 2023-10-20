package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is about Channel")

	/*

		Channel:- In Go, a channel is a powerful tool for communication and synchronization between goroutines (concurrent functions or threads).
		Think of a channel as a conduit that allows one goroutine to send data to another goroutine.
		Channels help coordinate and share data safely between concurrent parts of your program.

		Syntax :-
		 1:- var Channel_name chan type
		 2:- channel_name := make(chan type)

	*/

	firstchan := make(chan int)
	fmt.Println("value of the channel:- ", firstchan)
	fmt.Printf("Type of the value:- %T ", firstchan)
	fmt.Println()

	// home to put||insert or remove||get||recieve the value from channel

	// chan1 <- chan2  := or  chan1 <- value =  for sending = in this I am sending the value from chan2 to chan1  //  mtlb channel 2 se channel 1 me value transfer kar rhe ho
	// var  := <- chan1  = if we want to assign the channel in a varieble then we have to use this method
	// <- chan1 = if we want to print the value directly

	ch := make(chan int)
	fmt.Println("Hello from main")
	go multiplyWithChannel(ch)
	ch <- 10
	fmt.Println("Bye from main")

	///////////////////// close and check with ok syntax /////////////////////////////////

	ch1 := make(chan int)
	close(ch1)
	elem, ok := <-ch1 /// is command ko hamesha channel close ke bad hi run kre -agar ye ham channel close krne se phle try krege to deadlock error aayega mtlb ye wait krta rhta hai ki abhi koi function ya koi process hogi channel me channel close krne ke bad ye samaj jata hai ki channel close hai to erro nhi aati
	fmt.Println("Hello Go Guru Ji", ok, elem)

	// loops in channel

	c := make(chan string)
	go initStrings(c)

	for {
		resp, ok := <-c
		if ok == false {
			fmt.Println("Channel Close", ok)
			fmt.Println()
			break
		}
		fmt.Println("Channel Open", resp, ok)
	}

	// channel are two types :- 1:-Buffered= jin channel ki length hoti hai buffer channel hote hai
	//                          2:-Unbuffered channel = jin channel ki length nhi hoti

	mychnl := make(chan string, 9)
	mychnl <- "abc"
	mychnl <- "defghijklmnopqrst"
	mychnl <- "uvw"
	mychnl <- "xyz"

	fmt.Println("Length of the channel :=", len(mychnl))
	fmt.Println("Capacity of the channel :=", cap(mychnl))

	/// how to print the all elements of a channel = by using for range loop

	testch := make(chan string)

	go func() {
		testch <- "a"
		testch <- "b"
		testch <- "c"
		close(testch)
	}()

	for res := range testch {
		fmt.Println(res)
	}

}

/// these function are used in upper side of the code

func multiplyWithChannel(ch11 chan int) {
	fmt.Println(100 * <-ch11)
}

func initStrings(chnl chan string) {
	for v := 0; v < 3; v++ {
		chnl <- "Jai Ganpati Bappa"

	}
	close(chnl)
}

/* By Directional and Unidirectional Channel

By Directional = jin channel me ham ya to value dal sakte hai ya value nikal sakte hai dono kam ek hi channel se nhi kar skte hai

Unidirectional = jin channel se ham value nikal bhi sakte ho or value dal bhi sakte ho ek hi channel se ye dono kaam kar rhe ho

   Unidirectional hint = isme chan keyword ke bad ya to ya chan keyword ke phle (<-) ye symbol laga hota hai

   example :-  chan1 := make(chan <- string) || make(<- chan string)

   chan1 := make(chan <- string) ----> isme value keval bhej sakte hai recieve nhi kar sakte hai
   chan1 := make(<- chan string) ----> isme value keval recieve kar sakte hai value bhej nhi sakte hai


*/
