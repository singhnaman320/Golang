// [] Creating and accessing an Array []

// In Go language, arrays are created in two different ways:

// Using var keyword: In Go language, an array is created using the var keyword of a particular type with name, size, and elements. Syntax:

// Var array_name[length]Type

// Important Points:

// In Go language, arrays are mutable, so that you can use array[index] syntax to the left-hand side of the assignment to set the elements of
// the array at the given index.

// Var array_name[index] = element

// Approach 1: Using shorthand declaration:

// In Go language, arrays can also declare using shorthand declaration. It is more flexible than the above declaration.

// Syntax:

// array_name:= [length]Type{item1, item2, item3,...itemN}

// [] Multi-Dimensional Array

// As we already know that arrays are 1-D but you are allowed to create a multi-dimensional array. Multi-Dimensional arrays are the arrays of
// arrays of the same type. In Go language, you can create a multi-dimensional array using the following syntax:

// Array_name[Length1][Length2]..[LengthN]Type

// You can create a multidimensional array using Var keyword or using shorthand declaration as shown in the below example.

// Note: In a multi-dimension array, if a cell is not initialized with some value by the user, then it will initialize with zero by the compiler
// automatically. There is no uninitialized concept in the Golang.

// In an array, you can find the length of the array using len() method

package main

import "fmt"

func main() {

	fmt.Println("Learning arrays with golang..")

	// 1.
	var arrayOfCars [5]string

	arrayOfCars[0] = "Audi"
	arrayOfCars[1] = "BMW"
	arrayOfCars[2] = "Tata"
	arrayOfCars[3] = "Mahindra"
	arrayOfCars[4] = "Nisan"

	fmt.Println(arrayOfCars)
	fmt.Println("Length of arrayOfCars is", len(arrayOfCars))

	// 2.
	var missilesList = [5]string{"Bhramhos", "Prahar", "Prithvi", "Agni-V", "Nirbhay"}

	// OR missilesList := [5]string{"Bhramhos", "Prahar", "Prithvi", "Agni-V", "Nirbhay"}

	fmt.Println("Missiles list is", missilesList)

	fmt.Println("Length of missilesList is", len(missilesList))

	// 3. Multi-Dimensional

	// Creating and initializing
	// 2-dimensional array
	// Using shorthand declaration
	// Here the (,) Comma is necessary

	students := [3][3]string{{"Naman", "Kaushik", "Vishal"}, {"Nitesh", "Nikhil", "Prakhar"}, {"Nishant", "Prashant", "Adarsh"}}

	fmt.Println("Elements of multi- dimensional array-01 are: ")

	for i := 0; i < 3; i++ {

		var bag string = ""
		for j := 0; j < 3; j++ {

			bag = bag + students[i][j] + " "

		}
		fmt.Println(bag)
	}

	// 4.

	// Creating a 2-dimensional
	// array using var keyword
	// and initializing a multi
	// -dimensional array using index

	var numsList [2][2]int

	numsList[0][0] = 101
	numsList[0][1] = 201
	numsList[1][0] = 401
	numsList[1][1] = 104

	fmt.Println("Elements of multi- dimensional array-02 are: ")

	for k := 0; k < 2; k++ {

		for l := 0; l < 2; l++ {

			fmt.Println(numsList[k][l])

		}
	}
}
