/* By chatGpt


Methods:- n Go (often referred to as Golang), a method is a function associated with a type. Methods allow you to define behaviors that can be performed on values of a specific type. 
They are similar to functions, but they have a receiver,which is a parameter associated with the type they operate on.Methods are a way to add behaviors or operations to your custom data types.

Here's a simple explanation with an example:

Example 1-

package main

import "fmt"

// Define a struct named "Person" to represent a person's information.
type Person struct {
    FirstName string
    LastName  string
}

// Define a method called "FullName" for the "Person" type.
// This method takes a Person instance as a receiver and returns the full name as a string.
func (p Person) FullName() string {
    return p.FirstName + " " + p.LastName
}

func main() {
    // Create a Person instance
    person := Person{
        FirstName: "John",
        LastName:  "Doe",
    }

    // Call the "FullName" method on the "person" instance
    fullName := person.FullName()

    // Print the result
    fmt.Println("Full Name:", fullName)
}

In this example:

We define a struct type called "Person" with two fields: FirstName and LastName to represent a person's information.

We define a method named "FullName" for the "Person" type. This method takes a "Person" instance as a receiver (the p Person part) and returns the full name by concatenating the first name and last name.

In the main function, we create a "Person" instance called person with the first name "John" and last name "Doe."

We call the "FullName" method on the "person" instance using the dot notation (person.FullName()), and it returns the full name, which is then printed to the console.

Methods in Go are powerful because they allow you to attach behavior to your data structures, making your code more organized and easier to understand.







Example 2-



package main

import "fmt"

// Define a custom type Person
type Person struct {
    Name string
    Age  int
}

// Define a method for the Person type
func (p Person) SayHello() {
    fmt.Printf("Hello, my name is %s, and I'm %d years old.\n", p.Name, p.Age)
}

func main() {
    // Create a Person
    person1 := Person{
        Name: "Alice",
        Age:  30,
    }

    // Call the SayHello method
    person1.SayHello() // This is how you call a method on a value of the Person type
}



In this example:

We define a custom type Person with fields Name and Age.

We create a method named SayHello associated with the Person type. It takes a receiver parameter (p Person), which means it can be called on values of type Person.

Inside the SayHello method, we use the receiver p to access the fields of the Person value and print a greeting message.

In the main function, we create a Person named person1 and call the SayHello method on it. This method operates on the person1 value and prints a message specific to that person.

So, a method is essentially a function that is tied to a specific type and can access and operate on the data within values of that type. It's a way to add behaviors to your custom data types in Go.
*/

package main

import "fmt"

type data int

func (v1 data) div(v2 data) data {
	return v1 / v2
}

func main() {
	fmt.Println("This is about Methods")
	value1 := data(30)
	value2 := data(15)
	res := (value1).div(value2)
	fmt.Println(res)

}
