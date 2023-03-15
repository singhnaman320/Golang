package main

import "fmt"

func main() {

	fmt.Println("Learning if-else with Golang")

	// 1.
	num := 50 //:= is for declaration + assignment
	var outcome string

	if num > 50 {

		outcome = "More than half century completed.."

	} else if num == 50 {

		outcome = "Half century completed.."

	} else {

		outcome = "unable to complete half century.."
	}

	fmt.Println(outcome)

	// 2.

	checkNum := 18

	if checkNum%2 == 0 {

		fmt.Println("Given number is even")

	} else {

		fmt.Println("Given number is odd")
	}

	// 3.

	if num := 5; num < 12 {

		fmt.Println("Numbere is less than 12")

	} else {

		fmt.Println("Number is not in range")
	}

}
