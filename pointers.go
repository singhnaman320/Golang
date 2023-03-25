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

	var ptr *int = &value // Initialization of pointer ptr with memory address of variable value

	fmt.Println("Initialization of pointer ptr with memory address of variable value", ptr)

	// =-=-=-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=- [] Important Points =-=-=---=--=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

	// 1. The default value or zero-value of a pointer is always nil. Or you can say that an uninitialized pointer will always have a nil value.

	var intPointer *int
	var stringPointer *string

	fmt.Println("Integer pointer", intPointer)   // nil
	fmt.Println("String pointer", stringPointer) // nil

	// 2. Declaration and initialization of the pointers can be done into a single line.

	// (a.) var s *int = &a

	// 3. If you are specifying the data type along with the pointer declaration then the pointer will be able to handle the memory address of that
	// specified data type variable. For example, if you taking a pointer of string type then the address of the variable that you will give to a
	// pointer will be only of string data type variable, not any other type.

	// 4. To overcome the above mention problem you can use the Type Inference concept of the var keyword. There is no need to specify the data type
	// during the declaration. The type of a pointer variable can also be determined by the compiler like a normal variable. Here we will not use
	// the * operator. It will internally determine by the compiler as we are initializing the variable with the address of another variable.

	var string_value = "Learning_Pointers" // OR string_value := "Learning_Pointers"

	var valuePointer = &string_value // OR valuePointer := &string_value

	fmt.Println("String value is:", string_value)
	fmt.Println("Pointer(address) of string value is:", valuePointer)

	// [] Dereferencing the Pointer []

	// (* operator) is also termed as the dereferencing operator. It is not only used to declare the pointer variable but also used to access the
	// value stored in the variable which the pointer points to which is generally termed as indirecting or dereferencing. * operator is also termed
	// as the value at the address of

	fmt.Println("accessing string_value with pointer:", *valuePointer) //Learning_Pointers

	// We can also change the value of the pointer or at the memory location instead of assigning a new value to the variable.

	*valuePointer = "Learning Golang"
	fmt.Println("Value stored in string_value(*valuePointer) after Changing = ", string_value)

}
