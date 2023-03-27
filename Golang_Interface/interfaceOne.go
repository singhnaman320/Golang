package main

import "fmt"

// Go language interfaces are different from other languages. In Go language, the interface is a custom type that is used to specify a set of one
// or more method signatures and the interface is abstract, so you are not allowed to create an instance of the interface. But you are allowed to
// create a variable of an interface type and this variable can be assigned with a concrete type value that has the methods the interface requires.
// Or in other words, the interface is a collection of methods as well as it is a custom type.

// [ =-=-=-=-=-=] How to create an interface? [ =-=-=-=-=-=]

// In Go language, you can create an interface using the following syntax:

// type interface_name interface{

// // Method signatures

// }

// For Example:

// Creating an interface
// type myinterface interface{

// Methods
// fun1() int
// fun2() float64
// }

// Here the interface name is enclosed between the type and interface keywords and the method signatures enclosed in between the curly braces.

// [ =-=-=-=-=-=] How to implement interfaces? [ =-=-=-=-=-=]

// In the Go language, it is necessary to implement all the methods declared in the interface for implementing an interface. The go language
// interfaces are implemented implicitly. And it does not contain any specific keyword to implement an interface just like other languages as
// shown:

type firstInterface interface {
	area() float64
	perimeter() float64
}

type dimensions struct {
	length  float64
	breadth float64
}

// Implementing methods of the firstInterface interface
func (dim dimensions) area() float64 {

	return dim.length * dim.breadth
}

func (dim dimensions) perimeter() float64 {

	return 2 * (dim.length + dim.breadth)
}

func main() {

	// Accessing elements of the tank interface
	var inter firstInterface

	inter = dimensions{125, 25}

	fmt.Println("Area is:", inter.area())
	fmt.Println("Perimeter is:", inter.perimeter())
}


// [] Important Points []

// The zero value of the interface is nil.
// When an interface contains zero methods, such types of interface is known as the empty interface. So, all the types implement the empty 
// interface.

// Syntax:

// interface{}