package main

import "fmt"

// [=-=-=]Interface Types: The interface is of two types:

// [] one is static and
// [] another one is dynamic type.

// The static type is the interface itself, for example, secInterface in the below example. But interface does not have a static value so it
// always points to the dynamic values. A variable of the interface type containing the value of the Type which implements the interface, so the
// value of that Type is known as dynamic value and the type is the dynamic type. It is also known as concrete value and concrete type.

// Example:

type secInterface interface {
	area() float64
	perimeter() float64
}

func main() {

	var value secInterface

	fmt.Println("Value of the secInterface interface is: ", value)
	fmt.Printf("Type of the secInterface interface is: %T ", value)

}

// Output:

// Value of the secInterface interface is:  <nil>
// Type of the secInterface interface is: <nil>

// Here, in the above example, we have an interface named as a secInterface. In this example,
// fmt.Println(“Value of the secInterface interface is: “, value) statement returns the dynamic value of the interface and
// fmt.Printf(“Type of the secInterface interface is: %T “, value) statement returns the dynamic type, i.e nil because here the interface
// does not know who is implementing it.
