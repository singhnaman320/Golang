package main

import "fmt"

func main() {

	//9. (1)
	var vehicle string
	fmt.Println("taking Input.....")
	fmt.Scanln(&vehicle)
	fmt.Println("My car is " + vehicle)

	//9. (2)

	// the message we try here would be:
	// Naman got 8000 rupees in bank
	// or something follwoing the same format

	var name string
	var unit string
	var amount int
	var store string

	// taking input and storing in variable using the buffer string
	fmt.Scanln(&name, &store, &amount, &unit)

	// print out new string using the extracted values
	// witw your whole input taking in a single line only
	fmt.Printf("%d %s hold by %s in his %s\n", amount, unit, name, store)
}
