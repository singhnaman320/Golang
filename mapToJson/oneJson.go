package main

// encoding/json package provides utilities that can be used to convert to and from JSON. The same utility can be used to convert a golang map
// to JSON string and vice versa. A very important point to note though is that map allows integer values for keys while JSON doesn’t allow
// integer values for keys. JSON only allows string value for keys. So a map having an integer value for the key when converted to JSON will
// have a string value for the key.

import (
	"encoding/json"
	"fmt"
)

func main() {

	thisMap := make(map[int]string)

	thisMap[1] = "Naman"
	thisMap[2] = "Kaushik"

	value, err := json.Marshal(thisMap)

	if err != nil {

		fmt.Println("Error: %s", err.Error())

	} else {

		fmt.Println(string(value))
	}
}
