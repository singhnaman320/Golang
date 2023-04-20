package main

import (
	"fmt"
	"regexp"
)

// To store complicated regular expressions for reuse later, Compile() method parses a regular expression and returns a Regexp object if successful
// which can be used to match the text. Prototype of the function is:

// func Compile(expr string) (*Regexp, error)
// There are other various methods provided in the regexp package to match strings as shown:

func main() {

	// a regex object which can be reused later
	re, _ := regexp.Compile("regex")

	// string to be matched
	str := "learning regexp with Golang"

	checkMatch := re.FindStringIndex(str)

	fmt.Println(checkMatch) // Output: [9 14]
}
