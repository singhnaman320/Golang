package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Welcome to regular Expression")
	//regexp searching capacity
	sampleRegex := regexp.MustCompile("[.]")

	match := sampleRegex.Match([]byte("."))
	fmt.Println(match)

	//AND example of regexp
	first := "abc"
	second := "efg"
	third := "xyz"

	sRegex := regexp.MustCompile(first + second + third) //Ending operation
	match1 := sRegex.Match([]byte("abcefgxyz"))
	fmt.Println(match1)

	match1 = sRegex.Match([]byte("abc"))
	fmt.Println(match1)
	match1 = sRegex.Match([]byte("abcefg"))
	fmt.Println(match1)
	match1 = sRegex.Match([]byte("efg"))
	fmt.Println(match1)
	match1 = sRegex.Match([]byte("xyz"))
	fmt.Println(match1)

	//OR example of regexp

	// fourth:="abd"
	// fifth :="efg"

	sempleRegx := regexp.MustCompile("abd|efg")

	match2 := sempleRegx.Match([]byte("abd"))
	fmt.Println(match2)
	match2 = sempleRegx.Match([]byte("efg"))
	fmt.Println(match2)
	match2 = sempleRegx.Match([]byte("abdefg"))
	fmt.Println(match2)
	match2 = sempleRegx.Match([]byte("abc"))
	fmt.Println(match2)

	//Cart(^) and Dollar($) character
	sregexp1 := regexp.MustCompile("^abc$")
	match3 := sregexp1.Match([]byte("abcd"))
	fmt.Printf("for abcd: %t\n", match3)
	match3 = sregexp1.Match([]byte("1abc23"))
	fmt.Printf("for 1abc23: %t\n", match3)
	match3 = sregexp1.Match([]byte("abc"))
	fmt.Printf("for abc: %t\n", match3)

	//case sensitive example (?i)
	sampleRegex2 := regexp.MustCompile("(?i)abc")

	match4 := sampleRegex2.Match([]byte("abc"))
	fmt.Printf("For abc: %t\n", match4)

	match4 = sampleRegex2.Match([]byte("ABC"))
	fmt.Printf("For ABC: %t\n", match4)

}
