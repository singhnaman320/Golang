package main

import "fmt"

// [] Function Declaration

// Function declaration means a way to construct a function.
// Syntax:

// func function_name(Parameter-list)(Return_type){
//     // function body.....
// }

// The declaration of the function contains:

// func: It is a keyword in Go language, which is used to create a function.
// function_name: It is the name of the function.
// Parameter-list: It contains the name and the type of the function parameters.
// Return_type: It is optional and it contain the types of the values that function returns. If you are using return_type in your function,
// then it is necessary to use a return statement in your function.

// with return type

func rectanglePerimeter(len, bre int) int {

	perimeter := 2 * (len + bre)

	return perimeter
}

// [] without return type
func rectangleArea(length, breadth int) {

	area := length * breadth

	fmt.Println("Area of rectangle is:", area)
}

// [] Function Arguments

// In Go language, the parameters passed to a function are called actual parameters, whereas the parameters received by a function are called
// formal parameters.
// Note: By default Go language use call by value method to pass arguments in function.
// Go language supports two ways to pass arguments to your function:

// Call by value: In this parameter passing method, values of actual parameters are copied to function’s formal parameters and the two types of
// parameters are stored in different memory locations. So any changes made inside functions are not reflected in actual parameters of the caller.

func swapByValue(value1, value2 int) int {

	var num int
	num = value1
	value1 = value2
	value2 = num

	return num
}

// Call by reference: Both the actual and formal parameters refer to the same locations, so any changes made inside the function are actually
// reflected in actual parameters of the caller.

func swapByReference(val1, val2 *int) int {
	var store int
	store = *val1
	*val1 = *val2
	*val2 = store

	return store
}

// =--=-=--==-=-=-=-=-=--==- main =-=-==-=-=-=-==-=-=--=-=-=-=-=-=-=-=-=-

func main() {

	rectangleArea(6, 5)

	fmt.Println("Perimeter of rectangle is:", rectanglePerimeter(6, 5))

	// call by values
	var v1 int = 10
	var v2 int = 20
	fmt.Printf("\n v1 = %d and v2 = %d", v1, v2)

	swapByValue(v1, v2)
	fmt.Printf("\n v1 = %d and v2 = %d", v1, v2)

	// call by reference
	fmt.Println("By Reference")
	fmt.Printf("\nv1 = %d and v2 = %d", v1, v2)

	swapByReference(&v1, &v2)
	fmt.Printf("\np = %d and q = %d", v1, v2)
}

// =-=-=-=-=-==-=--==-=---=-=-=-=-=-=--=-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=-=
