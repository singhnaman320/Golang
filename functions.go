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

	perimeter := 2*(len + bre)

	return perimeter
}

// [] without return type
func rectangleArea(length, breadth int) {

	area := length * breadth

	fmt.Println("Area of rectangle is:", area)
}

// =--=-=--==-=-=-=-=-=--==- main =-=-==-=-=-=-==-=-=--=-=-=-=-=-=-=-=-=-

func main() {

	rectangleArea(6, 5)

	fmt.Println("Perimeter of rectangle is:", rectanglePerimeter(6, 5))
}

// =-=-=-=-=-==-=--==-=---=-=-=-=-=-=--=-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=-=
