package main

import "fmt"

// [] Using make function:

// You can also create a map using make() function. This function is an inbuilt function and in this method, you just need to pass the type
// of the map and it will return an initialized map.

// Syntax:

// make(map[Key_Type]Value_Type, initial_Capacity)
// make(map[Key_Type]Value_Type)

func main() {

	// Creating a map
    // Using make() function
	var mapByMake = make(map[float64]string)

	fmt.Println(mapByMake) // output : map[]

	 // As we already know that make() function always returns a map which is initialized So, we can add values in it

	mapByMake[12.5] = "Naman"
	mapByMake[17.9] = "Kaushik"
	mapByMake[10.5] = "Nishant"
	mapByMake[14.2] = "Prashant"
	mapByMake[16.7] = "Vishal"

	fmt.Println(mapByMake) // Output: map[10.5:Nishant 12.5:Naman 14.2:Prashant 16.7:Vishal 17.9:Kaushik]
}
