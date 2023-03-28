package main

import (
	"fmt"
)

// [] Type Switch:

// In Go interface, type switch is used to compare the concrete type of an interface with the multiple types provide in the case statements. It
// is similar to type assertion with only one difference, i.e, case specifies types, not the values. You can also compare a type to the interface
// type as shown below:

// Example:

func check(inter interface{}) {

	switch inter.(type) { // or switch val := inter(type) and print val for its value

	case int:
		fmt.Println("Type: int, Value:", inter.(int)) // use val

	case string:
		fmt.Println("\nType: string, Value:", inter.(string)) // use val

	case float64:
		fmt.Println("\nType: float64, Value:", inter.(float64)) // use val

	default:
		fmt.Println("\nUnknown type")

	}

}

func main() {

	check("Type Switch")
	check(58.54)
	check(false)
}
