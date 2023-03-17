// Go language contains only a single loop that is for-loop. A for loop is a repetition control structure that allows us to write a loop
// that is executed a specific number of times. In Go language, this for loop can be used in the different forms :

package main

import "fmt"

//=========================================================================================================================================

// 1. As simple for loop It is similar that we use in other programming languages like C, C++, Java, C#, etc.

//Syntax:

// for initialization; condition; post{
//        // statements....
// }

// The initialization statement is optional and executes before for loop starts. The initialization statement is always in a simple statement
// like variable declarations, increment or assignment statements, or function calls.

// The condition statement holds a boolean expression, which is evaluated at the starting of each iteration of the loop. If the value of the
// conditional statement is true, then the loop executes.

// The post statement is executed after the body of the for-loop. After the post statement, the condition statement evaluates again if the value
// of the conditional statement is false, then the loop ends.

// func main() {

// 	for i := 0; i < 5; i++ {

// 		fmt.Println("Learning simple for-loop..")
// 	}
// }

//=========================================================================================================================================

// 2. For loop as Infinite Loop: A for loop is also used as an infinite loop by removing all the three expressions from the for loop. When the
// user did not write condition statement in for loop it means the condition statement is true and the loop goes into an infinite loop.

// Syntax:

// for{
//      // Statement...
// }

// func main() {

// 	for {

// 		fmt.Println("Learning simple for-loop..")
// 	}
// }

//=========================================================================================================================================

// 3. for loop as while Loop: A for loop can also work as a while loop. This loop is executed until the given condition is true. When the
// value of the given condition is false the loop ends.

// Syntax:

// for condition{
//     // statement..
// }

// func main() {

// 	num := 5

// 	for num < 20 {

// 		num += num
// 	}
// 	fmt.Println(num)
// }

//=========================================================================================================================================

// 4. Simple range in for loop: You can also use the range in the for loop.

// Syntax:

// for i, j:= range rvariable{
//    // statement..
// }
// Here,

// i and j are the variables in which the values of the iteration are assigned. They are also known as iteration variables.
// The second variable, i.e, j is optional.
// The range expression is evaluated once before the starting of the loop.

// func main() {

// 	// Here, we haven taken range variable as array
// 	rvariable := []string{"simple", "range", "for", "loops"}

// 	// i and j stores the value of rvariable
// 	// i store index number of individual string and
// 	// j store individual string of the given array

// 	for i, j := range rvariable {

// 		fmt.Println(i, j)
// 	}
// }

//=========================================================================================================================================

// 5. Using for loop for strings: A for loop can iterate over the Unicode code point for a string.

// Syntax:

// for index, chr:= range str{
//      // Statement..
// }
// Here, The index is the variable which store the first byte of UTF-8 encoded code point and chr store the characters of the given string and
// str is a string.

// func main() {

// 	// String as a range in the for loop
// 	for index, j := range "aXbZcY" {

// 		fmt.Printf("The index number of %U is %d\n", j, index)
// 	}
// }

//=========================================================================================================================================

// 6. For Maps: A for loop can iterate over the key and value pairs of the map.

// Syntax:

// for key, value := range map {
//      // Statement..
// }

// func main() {

// 	mapping := map[int]string{
// 		10: "Learning",
// 		20: "Loops",
// 		30: "With",
// 		40: "Golang",
// 	}

// 	for key, value := range mapping {

// 		fmt.Println(key, "-->", value)
// 	}
// }

//=========================================================================================================================================

// 7. For Channel: A for loop can iterate over the sequential values sent on the channel until it closed.

// Syntax:

// for item := range Chnl {
//      // statements..
// }

// func main() {

// 	// using channel
// 	chnl := make(chan int)
// 	go func() {
// 		chnl <- 100
// 		chnl <- 1000
// 		chnl <- 10000
// 		chnl <- 100000
// 		close(chnl)
// 	}()

// 	for value := range chnl {

// 		fmt.Println(value)
// 	}
// }

//=========================================================================================================================================

// All in one

func main() {

	//1.
	for i := 0; i < 5; i++ {

		fmt.Println("Learning simple for-loop..")
	}

	//2.
	for {

		fmt.Println("Learning simple for-loop..")
	}

	//3.
	num := 5

	for num < 20 {

		num += num
	}
	fmt.Println(num)

	//4.
	rvariable := []string{"simple", "range", "for", "loops"}

	// i and j stores the value of rvariable
	// i store index number of individual string and
	// j store individual string of the given array

	for i, j := range rvariable {

		fmt.Println(i, j)
	}

	//5.
	for index, j := range "aXbZcY" {

		fmt.Printf("The index number of %U is %d\n", j, index)
	}

	//6.
	mapping := map[int]string{
		10: "Learning",
		20: "Loops",
		30: "With",
		40: "Golang",
	}

	for key, value := range mapping {

		fmt.Println(key, "-->", value)
	}

	//7. using channel
	chnl := make(chan int)
	go func() {
		chnl <- 100
		chnl <- 1000
		chnl <- 10000
		chnl <- 100000
		close(chnl)
	}()

	for value := range chnl {

		fmt.Println(value)
	}
}
