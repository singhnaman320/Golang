package main

import "fmt"

// [] Method Can Accept both Pointer and Value:

// As we know that in Go, when a function has a value argument, then it will only accept the values of the parameter, and if you try to pass a
// pointer to a value function, then it will not accept and vice versa. But a Go method can accept both value and pointer, whether it is defined
// with pointer or value receiver.

// Employee structure
type employee struct {
	name       string
	department string
}

// Method with a pointer receiver of employee type
func (emp1 *employee) firstEmployeeDetails(dept string) {

	(*emp1).department = dept
}

// Method with a value receiver of employee type
func (emp2 employee) secEmployeeDetails() {

	emp2.name = "Naman"
	fmt.Println("Name of employee before is:", emp2.name)
}

func main() {

	// Initializing the values of the employee structure
	result := employee{

		name:       "Kaushik",
		department: "Human Resource",
	}

	fmt.Println("Name of employee department before is:", result.department)

	// Calling the firstEmployeeDetails method (pointer method) with value
	result.firstEmployeeDetails("Engineering")
	fmt.Println("Name of employee department after is:", result.department)

	// Calling the secEmployeeDetails method (value method) with a pointer
	(&result).secEmployeeDetails()
	fmt.Println("Name of employee after is:", result.name)

}
