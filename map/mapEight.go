package main

import "fmt"

// Modifying map:

// As we know that maps are of reference type. So, when we assign an existing map to a new variable, both the maps still refer to the same
// underlying data structure. So, when we update one map it will reflect in another map.

func main() {

	// Creating and initializing a map
	myMap := map[int]string{

		10: "Naman",
		12: "Kaushik",
		30: "Vishal",
		14: "Nishant",
		9:  "Prashant",
	}

	fmt.Println("\nOriginal map: ", myMap)

	// Assigned the map into a new variable
	newMap := myMap

	// Perform modification in new_map
	newMap[30] = "Maharana"
	newMap[9] = "Shivaji"

	// Display after modification
	fmt.Println("\nNew map: ", myMap)

	fmt.Println("\nModification done in old map:\n", myMap)

}
