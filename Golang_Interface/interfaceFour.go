package main

import "fmt"

// value, ok := a.(T)

// Here if the type of the a is equal to T, then the value contains the dynamic value of the a and ok will set to true. And if the type of the
// a is not equal to T, then ok set to false and value contain zero value, and the program does not panic as shown below:

func check(inter interface{}) {

	value, ok := inter.(float64)

	fmt.Println(value, ok)
}

func main() {

	var val1 interface{} = 66.7

	check(val1) // 66.7 true

	var val2 interface{} = "Golang"

	check(val2) // 0 false
}
