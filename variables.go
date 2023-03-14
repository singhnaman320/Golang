package main

import "fmt"

// Global variable
const globalVariable string = "I am Global variable"

func main() {

	//Declaring a variable ----> var variable VariableType
	//Note: Unlike in other languages variable declared in go language must be used otherwise it will give compile time error

	var num int = 10

	// Or var num int;
	// num= 10;

	fmt.Println(num) // If any value is not given, it will print Output= 0 similarly for string default value is null;

	//=====================================================================================

	// The fmt.out.printf() function in Golang allows users to print formatted data.
	// The function takes a template string that contains the text that needs to be formatted, plus some annotation verbs that tell the
	// fmt functions how to format the trailing arguments.

	// Conversion characters
	// Conversion characters tell Golang how to format different data types. Some of the most commonly used specifiers are:

	// v – formats the value in a default format
	// d – formats decimal integers
	// g – formats the floating-point numbers
	// b – formats base 2 numbers
	// o – formats base 8 numbers
	// t – formats true or false values
	// s – formats string values

	//1.
	var str string = "Naman Singh"
	fmt.Printf("The string is %s \n", str)

	//2.
	var data int = 58
	fmt.Printf("The decimal value is %d \n", data) //If we not use %d then it will give compile time error

	//3.
	// float64 is default type of a float value in Go and it takes 64 bit
	//A variable of type float64 can store decimal numbers ranging from 2.2E-308 to 1.7E+308. This applies to negative and positive numbers.

	// It is also possible to create const values of type float64 instead of variables. The only difference is that const values are just
	// variables whose values can not be changed from what it was initialized to.

	//A variable of type float32 can store decimal numbers ranging from 1.2E-38 to 3.4E+38. This is the range of magnitudes that
	// a float32 variable can store. This applies to negative and positive numbers.

	// If we want to get variable type we use [ %T ] Ex: fmt.Printf("The floating point is %g \n", floatValue0)

	var floatValue0 float64 = 3.142
	fmt.Printf("The floating point is %g \n", floatValue0)

	var floatValue1 float32 = 3.142
	fmt.Printf("The floating point is %g \n", floatValue1)

	//4.
	var baseOf2 int = 18
	fmt.Printf("The binary value of baseOf2 is %b \n", baseOf2)

	//5.
	var baseOf8 int = 18
	fmt.Printf("The value of baseOf8 is %o \n", baseOf8)

	//6.
	var check1 int = 18
	var check2 int = 18
	fmt.Printf("The boolean value is %t \n", check1 == check2)

	//7. Lets see boolean value
	var booleanValue bool = true
	fmt.Println(booleanValue)

	//8. implicit type -> with no datatype
	var website = "naman.goole.com"
	fmt.Println(website)

	users := 2800
	fmt.Println(users)

	allUsers := 2800.0
	fmt.Println(allUsers) // also give int value as output i.e, 2800

	//accessing global variable

	fmt.Println(globalVariable)
	fmt.Printf("Datatype of global variable is: %T \n", globalVariable)

}
