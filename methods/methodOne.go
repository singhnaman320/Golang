package main

import "fmt"

// Go language support methods. Go methods are similar to Go function with one difference, i.e, the method contains a receiver argument in it.
// With the help of the receiver argument, the method can access the properties of the receiver. Here, the receiver can be of struct type or
// non-struct type. When you create a method in your code the receiver and receiver type must be present in the same package. And you are not
// allowed to create a method in which the receiver type is already defined in another package including inbuilt type like int, string, etc. If
// you try to do so, then the compiler will give an error.

// [] Syntax:

// func(reciver_name Type) method_name(parameter_list)(return_type){
// // Code
// }
// Here, the receiver can be accessed within the method.

// [] struct: [] In Go programming, a structure or struct is a user-defined type to store a collection of different fields into a single field.

// Method with struct type receiver
// In Go language, you are allowed to define a method whose receiver is of a struct type. This receiver is accessible inside the method as shown:

// Employee structure (will be using as reciever)

type employee struct {
	name       string
	empId      int
	department string
	salary     int
}

// Method with a receiver of employee type
func (givenEmployee employee) details() {

	fmt.Println("Name of employee is:", givenEmployee.name)
	fmt.Println("Employee Id is:", givenEmployee.empId)
	fmt.Println("Employee department is:", givenEmployee.department)
	fmt.Println("Employee salary is:", givenEmployee.salary)
}
func main() {

	// Initializing the values of the employee structure

	output := employee{

		name:       "Naman",
		empId:      23,
		department: "Engineeering",
		salary:     60000,
	}

	output.details()
}
