package main

import (
	"fmt"
)

// [] Methods with Pointer Receiver:

// In Go language, you are allowed to create a method with a pointer receiver. With the help of a pointer receiver, if a change is made in the
// method, it will reflect in the caller which is not possible with the value receiver methods.

// Syntax:

// func (p *Type) method_name(...Type) Type {
// 	  Code
// }

type employee struct {
	name       string
	department string
}

// Method with a receiver of employee type
func (emp *employee) employeeDetails(dept string) {

	(*emp).department = dept // OR emp.department = dept
}

func main() {

	// Initializing the values of the author structure
	result := employee{

		name:       "Naman",
		department: "Engineering",
	}

	fmt.Println("Employee's name:", result.name)
	fmt.Println("Employee's department:", result.department)

	// Creating a pointer
	net := &result
	net.employeeDetails("Testing")

	// Calling the show method

	//fmt.Println("Employee's name:", net.name)
	//fmt.Println("Employee's department:", net.department)  OR

	fmt.Println("Employee's name:", result.name)
	fmt.Println("Employee's department:", result.department)

}
