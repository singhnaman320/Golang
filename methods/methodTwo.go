package main

import "fmt"

// [] Method with Non-Struct Type Receiver:

// In Go language, you are allowed to create a method with non-struct type receiver as long as the type and the method definitions are present in
// the same package. If they present in different packages like int, string, etc, then the compiler will give an error because they are defined in
// different packages.

// Type definition
type data int

func (value1 data) multiply(value2 data) data {

	return value1 * value2
}

// [=-=-=-=--=-] NOTE [=-=-=-=-=]
/*
// if you try to run this code,
// then compiler will throw an error
func(value1 int)multiply(value2 int)int{
return d1 * d2
}
*/

func main() {

	num1 := data(25)
	num2 := data(25)

	net := num1.multiply(num2)

	fmt.Println("final output is:", net)
}
