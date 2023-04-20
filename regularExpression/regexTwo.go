package main

import (
	"fmt"
	"regexp"
	"strings"
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

	checkMatchOne := re.FindStringIndex(str)

	fmt.Println(checkMatchOne) // Output: [9 14]

	//=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=--=-=-=-=-=-=-=-=-=--==-=-=-=-=-=-=-=-=-=-=-=

	str1 := "Learning with Golang"

	// prints an empty slice as there is no match
	checkMatchTwo := re.FindStringIndex(str1)

	fmt.Println(checkMatchTwo)

	//=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=--=-=-=-=-=-=-=-=-=--==-=-=-=-=-=-=-=-=-=-=-=

	// finds the first or leftmost match to a given pattern.
	res, _ := regexp.Compile("[0-9]+-n.*g")

	// matches one or more numbers followed by n and any number of characters upto g
	checkMatchThree := res.FindString("2024-naman_singh")

	fmt.Println(checkMatchThree) // Output: 2024-naman_sing

	//=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=--=-=-=-=-=-=-=-=-=--==-=-=-=-=-=-=-=-=-=-=-=

	// returns a slice of all successive matches of the expression
	checkMatchFour := re.FindAllStringSubmatchIndex("learning regexp with Golang from regexp to regex", -1)

	fmt.Println(checkMatchFour) // Output: [[9 14] [33 38] [43 48]]

	//=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=--=-=-=-=-=-=-=-=-=--==-=-=-=-=-=-=-=-=-=-=-=

	// returns a copy and replaces matches with the replacement string
	re2, _ := regexp.Compile(" ")

	checkMatchFive := re2.ReplaceAllString("Learning regexp with Golang", "-")

	fmt.Println(checkMatchFive) // Output: Learning-regexp-with-Golang

	//=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=--=-=-=-=-=-=-=-=-=--==-=-=-=-=-=-=-=-=-=-=-=

	// returns a copy in which all matches are replaced by return value of function
	re3, _ := regexp.Compile("[aeiou]")

	checkMatchSix := re3.ReplaceAllStringFunc("All I do is code everytime.", strings.ToUpper)

	fmt.Println(checkMatchSix) // Output: All I dO Is cOdE EvErytImE.

	//=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=--=-=-=-=-=-=-=-=-=--==-=-=-=-=-=-=-=-=-=-=-=
}
