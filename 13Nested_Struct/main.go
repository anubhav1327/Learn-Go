package main

import "fmt"

type user struct {
	name    string
	age     int
	address address
}

type address struct {
	name1    string
	age1     int
	address1 string
}

func main() {
	fmt.Println("This is about Nested_Struct")
	var u user
	fmt.Println(u)
	u = user{"anu", 23, address{"singh", 24, "hardoi"}}
	fmt.Println(u)
	u = user{
		name: "anu1",
		age:  231,
		address: address{
			name1:    "Singh1",
			age1:     241,
			address1: "Hardoi",
		},
	}
	fmt.Println(u)
	fmt.Println(u.name)
	fmt.Println(u.address)
	fmt.Println(u.address.name1)
	fmt.Println(u.address.age1)

	//Anonymous Struct = isme dikkat ye hai ki agar isme koi value dege to vo ek hi bar de sakte hai jaise agar ek bar string assign kr diya to dubara nhi assign kr sakte hai to is error ko dur krne ke liye ham isi me name assign kr dete hai jaise second string ya first string second int ya first int jaha first or second as a name work krege string me isiliye isme name assign karte hai is error ko dur karn ke liye ///////////////////////////////

	test := struct {
		int
		string
	}{
		123, "test",
	}

	fmt.Println(test)

	test2 := struct {
		first   int
		second  int
		first1  string
		second1 string
	}{
		123, 334, "test", "radhe",
	}

	fmt.Println(test2)
}
