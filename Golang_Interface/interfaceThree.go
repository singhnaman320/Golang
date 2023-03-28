package main

import "fmt"

// [] Type Assertions:
// In Go language, type assertion is an operation applied to the value of the interface. Or in other words, type assertion is a process to extract
// the values of the interface.

// Syntax:
// a.(T)

// Here, a is the value or the expression of the interface and T is the type also known as asserted type. The type assertion is used to check that
// the dynamic type of its operand will match the asserted type or not. If the T is of concrete type, then the type assertion checks the given
// dynamic type of a is equal to the T, here if the checking proceeds successfully, then the type assertion returns the dynamic value of a. Or if
// the checking fails, then the operation will panics. If the T is of an interface type, then the type assertion checks the given dynamic type of
// a satisfies T, here if the checking proceeds successfully, then the dynamic value is not extracted.

// Example:

func check(inter interface{}) {

	val := inter.(string) // string is type

	fmt.Println("Value is:", val)

}

func main() {

	var interfaceValue interface{} = "Interface Golang"

	check(interfaceValue)
}

// In the above example if we change this val := a.(string) statement into val := a.(int), then the program panic. So to overcome this problem we
// use the following syntax:

// value, ok := a.(T)

// we'll see it in interfaceFour.go
