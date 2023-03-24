package main

import "fmt"

// Pointers in Go programming language or Golang is a variable that is used to store the memory address of another variable. Pointers in Golang
// is also termed as the special variables. The variables are used to store some data at a particular memory address in the system. The memory
// address is always found in hexadecimal format(starting with 0x like 0xFFAAF etc.).

func main() {

	// 1.

	// In this program, we are storing the hexadecimal number into a variable. But you can see that the type of values is int and saved as the
	// decimal or you can say the decimal value of type int is storing. But the main point to explain this example is that we are storing a hexadecimal
	// value(consider it a memory address) but it is not a pointer as it is not pointing to any other memory location of another variable. It is just a
	// user-defined variable. So this arises the need for pointers.

	// storing the hexadecimal values in

	value1 := 0xFF
	value2 := 0x9C

	fmt.Printf("Type of variable value1 is %T\n", value1)
	fmt.Printf("Value of value1 in hexadecimal is %X\n", value1)
	fmt.Printf("Value of value1 in decimal is %v\n", value1)

	fmt.Printf("Type of variable value2 is %T\n", value2)
	fmt.Printf("Value of value2 in hexadecimal is %X\n", value2)
	fmt.Printf("Value of value2 in decimal is %v\n", value2)

	// 2.

	// A pointer is a special kind of variable that is not only used to store the memory addresses of other variables but also points where the
	// memory is located and provides ways to find out the value stored at that memory location. It is generally termed as a Special kind of Variable
	// because it is almost declared as a variable but with *(dereferencing operator).

	// [] Declaration and Initialization of Pointers []

	// (* Operator) also termed as the dereferencing operator used to declare pointer variable and access the value stored in the address.
	// (& operator) termed as address operator used to returns the address of a variable or to access the address of a variable to a pointer.

	// [] Declaring a pointer:

	// var pointer_name *Data_Type

	// Example: Below is a pointer of type string which can store only the memory addresses of string variables.

	// var s *string

	// Initialization of Pointer: To do this you need to initialize a pointer with the memory address of another variable using the address
	// operator

	var value = 20 // normal variable declaration

	var ptr *int = &value  // Initialization of pointer ptr with memory address of variable value

	fmt.Println("Initialization of pointer ptr with memory address of variable value", ptr)
}
