package main

import "fmt"

// [] Iteration over a map:

// You can iterate a map using the range for loop. The value of this loop may vary because the map is an unordered collection.

func main() {

	mapIteration := map[int]string{

		10: "Naman",
		12: "Kaushik",
		30: "Vishal",
		14: "Nishant",
		9:  "Prashant",
	}

	for roll, studentName := range mapIteration {

		fmt.Println(roll, studentName)
	}
}
