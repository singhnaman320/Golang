package main

import (
	"fmt"
	"regexp"
)

// A Regular Expression (or RegEx) is a special sequence of characters that defines a search pattern that is used for matching specific text.
// In Golang, there’s a built-in package for regular expressions, called the regexp package which contains all list of actions like filtering,
// replacing, validating, or extracting. It uses the RE2 syntax standards. The MatchString() function reports whether the string passed as a
// parameter contains any match of the regular expression pattern.

// Syntax:

// func MatchString(pattern string, s string)

// Returns: matched bool, err error

func main() {

	// string in which the pattern is to be found
	givenStringOne := "Pneumonoultramicroscopicsilicovolcanoconiosis"

	// returns true if the pattern is present in the string, else returns false err is nil if the regexp is valid
	checkMatchOne, err := regexp.MatchString("microscopic", givenStringOne)

	fmt.Println("Check Match: ", checkMatchOne, "Check Error: ", err)

	// this returns false as the match is unsuccessful
	givenStringTwo := "regrexp with Golang"

	checkMatchTwo, erro := regexp.MatchString("naman", givenStringTwo)

	fmt.Println("Check Match: ", checkMatchTwo, "Check Error: ", erro)

	// this will throw an error as the pattern is not valid
	checkMatchThree, er := regexp.MatchString("nam(an", givenStringTwo)

	fmt.Println("Check Match: ", checkMatchThree, "Check Error: ", er)
}
