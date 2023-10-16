/*
By chatGpt:-

Struct :- In Go (or Golang), a struct is a composite data type that groups together variables with different data types under a single name. Think of it like a container that can hold various pieces of data, much like a struct in the real world can have different properties like height, weight, and color.

Here's a simple explanation of a struct with an example:

// Define a struct
type Person struct {
    Name    string
    Age     int
    IsMale  bool
}

func main() {
    // Create a variable of type 'Person'
    person1 := Person{
        Name:   "Alice",
        Age:    30,
        IsMale: false,
    }

    // Access and modify struct fields
    fmt.Println("Name:", person1.Name)
    fmt.Println("Age:", person1.Age)

    // Update a struct field
    person1.Age = 31

    // Print the updated struct
    fmt.Println("Updated Age:", person1.Age)
}


In this example, we define a Person struct with three fields: Name, Age, and IsMale. We then create an instance of the Person struct called person1 and initialize it with specific values.

You can access and modify the fields of the person1 struct just like you would with any other variable. This allows you to group related data together in a structured way.

In summary, a struct in Go is a custom data type that lets you bundle multiple variables together under a single name, and it's commonly used to represent objects or data with different attributes in your programs.





*/

package main

import "fmt"

type user struct {
	name     string
	age      int
	address  string
	address2 interface{}
}

func main() {
	fmt.Println("This is about Structure or struct")
	var u user
	fmt.Println(u)
	u = user{"Anubhav Sing", 23, "Hardoi", nil}
	fmt.Println(u)
	u = user{
		name:     "Anubhav Singh",
		age:      23,
		address:  "hardoi 2",
		address2: nil,
	}

	fmt.Println(u)
	fmt.Println(u.name)
	fmt.Println(u.age)
}
