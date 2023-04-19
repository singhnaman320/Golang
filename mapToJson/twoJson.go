package main

import (
	"encoding/json"
	"fmt"
)

// Example where we convert a map to a JSON where we have a struct for the value in the map. Below is the code for that

type employee struct {
	Name string
}

func main() {

	myMap := make(map[int]employee)

	myMap[1] = employee{Name: "Naman"}

	value, err := json.Marshal(myMap)

	if err != nil {

		fmt.Printf("Error: %s", err.Error())

	} else {

		fmt.Println(string(value))
	}

}

// Output: {"1":{"Name":"Naman"}}
