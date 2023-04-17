package main

import "fmt"

// Golang Maps is a collection of unordered pairs of key-value. It is widely used because it provides fast lookups and values that can retrieve,
// update or delete with the help of keys.

// How to creating and initializing Maps?

// In Go language, maps can create and initialize using two different ways:

// 1. Simple: In this method, you can create and initialize a map without the use of make() function:

// Creating Map: You can simply create a map using the given syntax:

// An Empty map
// map[Key_Type]Value_Type{}

// Map with key-value pair
// map[Key_Type]Value_Type{key1: value1, ..., keyN: valueN}

// Example:

// var mymap map[int]string

// In maps, the zero value of the map is nil and a nil map doesn’t contain any key. If you try to add a key-value pair in the nil map, then
// the compiler will throw runtime error.

// Initializing map using map literals: Map literal is the easiest way to initialize a map with data just simply separate the key-value pair with
// a colon and the last trailing colon is necessary if you do not use, then the compiler will give an error.

func main() {

	// Creating and initializing empty map using var keyword
	var firstMap map[int]string

	// Checking if the map is nil or not
	if firstMap == nil {

		fmt.Println("True")

	} else {

		fmt.Println("False")
	}

	// Creating and initializing a map using shorthand declaration and using map literals
	secMap := map[int]string{

		10: "This is 10",
		20: "This is 20",
		30: "This is 30",
		40: "This is 40",
		50: "This is 50",
	}

	fmt.Println("secMap is:", secMap)
}
