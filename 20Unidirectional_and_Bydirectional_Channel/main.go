/* By Directional and Unidirectional Channel

By Directional = jin channel me ham ya to value dal sakte hai ya value nikal sakte hai dono kam ek hi channel se nhi kar skte hai

Unidirectional = jin channel se ham value nikal bhi sakte ho or value dal bhi sakte ho ek hi channel se ye dono kaam kar rhe ho

   Unidirectional hint = isme chan keyword ke bad ya to ya chan keyword ke phle (<-) ye symbol laga hota hai

   example :-  chan1 := make(chan <- string) || make(<- chan string)

   chan1 := make(chan <- string) ----> isme value keval bhej sakte hai recieve nhi kar sakte hai
   chan1 := make(<- chan string) ----> isme value keval recieve kar sakte hai value bhej nhi sakte hai


*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("TITLE FOR THIS CODE :- This is about Unidirectional and Bydirectinal Channel")

	// creating a bidirectional channel

	chan1 := make(chan string)
	chan2 := make(chan string)

	go sending(chan1)
	valueFromChannel := <-chan1
	fmt.Println("valueFromChannel:- ", valueFromChannel)
	go recieving(chan2)
	chan2 <- valueFromChannel // for this line go 5.00 min in video no 30

	// creating a unidirectional channel on for recieving

	chan3 := make(chan string)
	go convert(chan3)
	fmt.Println(chan3)
	fmt.Println(<-chan3)
}

// creatiing a function for unidirectional channel : - in this function we are recieving a channel and the channel is working on for send only

func convert(s chan string) {
	s <- "from conver function unidirectional channal :-Ganpati Bappa"
}

// creating a function for sending the value in channel

func sending(s chan string) {
	s <- "Ganpati Bappa"
}

// // creating a function for recieving the value in channel

func recieving(s chan string) {
	fmt.Println("getting value :-", <-s)
}
