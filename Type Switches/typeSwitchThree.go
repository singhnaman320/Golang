package main

import "fmt"

func find_type(value interface{}) {

	switch v_type := value.(type) {

	case int, float64:

		// type is an int or float
		fmt.Println("Type is a number, either an int or a float:", v_type)

	case string, bool:

		// type is a string or bool
		fmt.Println("Type is either a string or a bool:", v_type)

	case *int, *bool:

		// type is either an int pointer or a bool pointer
		fmt.Println("Type is a pointer to an int or a bool:", v_type)

	default:

		// type of the interface is unknown
		fmt.Println("Type unknown!")
	}
}
func main() {

	// an interface that has a string value
	var value interface{} = "type switch"

	// calling the find_type method to determine type of interface
	find_type(value)

	// re-assigning v with a float64 value
	value = 87.65

	find_type(value)

}
