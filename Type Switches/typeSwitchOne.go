package main

import "fmt"

// A switch is a multi-way branch statement used in place of multiple if-else statements but can also be used to find out the dynamic type of an
// interface variable. A type switch is a construct that performs multiple type assertions to determine the type of variable (rather than values)
// and runs the first matching switch case of the specified type. It is used when we do not know what the interface{} type could be.

func main() {

	// an interface that has a string value

	var value interface{} = "Type Switches"

	// type switch to find out interface{} type

	switch interfaceType := value.(type) {

	case int64:
		fmt.Println("Type is a integer:", interfaceType)

	case float64:
		fmt.Println("Type is a float:", interfaceType)

	case string:
		fmt.Println("Type is a string:", interfaceType)

	case nil:
		fmt.Println("Type is a nil")

	case bool:
		fmt.Println("Type is a boolean:", interfaceType)

	default:
		fmt.Print("Unknown type")
	}
}
