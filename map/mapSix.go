package main

import "fmt"

// Check the existence of the key in the map:

// In maps, you can check whether the given key exists or not using the following syntax:

// With value
// It will gives the value and check result
// value, check_variable_name:= map_name[key]

// or

// Without value using the blank identifier
// It will only give check result
// _, check_variable_name:= map_name[key]

// Here, if the value of the check_variable_name is true which means the key exists in the given map and if the value of check_variable_name
// is false which means the key does not exist in the given map.

func main() {

	myMap := map[int]string{

		10: "Naman",
		12: "Kaushik",
		30: "Vishal",
		14: "Nishant",
		9:  "Prashant",
	}

	// Checking the key is available
	// or not in the myMap map
	studentName, ok := myMap[30]

	fmt.Println("Is key present :", ok)
	fmt.Println("Value hold by partcular key: ", studentName)

	// Using blank identifier
	_, ok1 := myMap[9]
	fmt.Println("Is key present :", ok1)

}
