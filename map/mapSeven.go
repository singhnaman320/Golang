package main

import "fmt"

// Delete key from the map:

// In maps, you are allowed to delete the key present in the map using the delete() function. It is inbuilt function and does not return any
// value and does not do anything if the key does not present in the given map. In this function, you just simply pass the map and key which
// you want to delete from the map.

// Syntax:

// delete(map_name, key)

func main() {

	myMap := map[int]string{

		10: "Naman",
		12: "Kaushik",
		30: "Vishal",
		14: "Nishant",
		9:  "Prashant",
	}

	fmt.Println("original map is: ", myMap)

	// Deleting keys using delete function

	delete(myMap, 9)
	delete(myMap, 30)

	fmt.Println("Map after deletion is: ", myMap)
}
