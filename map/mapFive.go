package main

import "fmt"

// Retrieve a value related to a key in the maps?:

// In maps, you can retrieve a value with the help of key using the following syntax:

// map_name[key]

// If the key doesn’t exist in the given map, then it will return zero value of the map, i.e, nil. And if the key exists in the given map, then
// it will return the value related to that key.

func main() {

	thisMap := map[int]string{

		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}

	firstValue := thisMap[93]
	secValue := thisMap[91]

	fmt.Println("Value of key[93]: ", firstValue)
	fmt.Println("Value of key[91]: ", secValue)
}
