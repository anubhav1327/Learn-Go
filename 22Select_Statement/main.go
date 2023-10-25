package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("This is about Select Statement")

	/*

		select statement == switch case me kya hota hai ki jo condition case se match kar rhi hogi vo output print hoga
		                    otherwise default case exceute ho jayega lekin SELECT STATEMENT  me aisa nhi hota hai isme cases
							communication ko reffer krte hai (send or recieve case ko reffer krte hai) mtlb send or recieve kr according
							koi operation ho rha hoga to output de dege wrna default case execute ho jayega

		syntax - select{
			         case Send || Receive case   // statement

					 .......
		             default:                    // statement
		         }

	*/

	chnl1 := make(chan string)
	chnl2 := make(chan int)

	go functionOne(chnl1) // dono go routine run hogi lekin output keval first
	// wali goroutine ka aayega =  Welcome to channel 1 with Ganpti bappa kyuki aisa isliye hoga ki jo case phle aayega wahi phle execute hoga jo case bad
	//  me aayega vo bad me execute hoga
	go functionTwo(chnl2)

	select {

	// agar select statement ke andar koi statement nhi likhege to ye program ko break kr dega or DEADLOCK error dega isliye agar kuch na likhne ki gajha ham default statement likhte hai
	//  example for default statement := select {
	//      default :
	// 	 fmt.Println("Jai Ganpati Bappa Maurya")
	// }
	case val1 := <-chnl1:
		fmt.Println(val1)
	case val2 := <-chnl2:
		fmt.Println(val2)
	}
}

func functionOne(chnl1 chan string) {
	time.Sleep(1 * time.Second)
	chnl1 <- "Welcome to channel 1 with Ganpati bappa"
}

func functionTwo(chnl2 chan int) {
	time.Sleep(1 * time.Second)
	chnl2 <- 20
}
