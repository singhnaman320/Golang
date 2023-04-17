package main

import "fmt"

// Add key-value pairs in the map:

// In maps, you are allowed to add key-value pairs in the initialized map using the given syntax:

// map_name[key]=value

// In maps, if you try to add an already existing key, then it will simply override or update the value of that key with the new value.

func main() {

	// Creating and initializing a map
	myMap := map[int]string{

		10: "Naman",
		12: "Kaushik",
		30: "Vishal",
		14: "Nishant",
		9:  "Prashant",
	}

	fmt.Println("Original map:", myMap)

	// Adding new key-value pairs in the map
	myMap[60] = "Rishabh"
	myMap[22] = "Abhinav"

	fmt.Println("\nMap after adding new key-value pair:\n", myMap)

	// Updating values of the map
	myMap[30] = "Tushar"
	myMap[9] = "Yashika"

	fmt.Println("\nMap after updating values of the map:\n", myMap)
}
