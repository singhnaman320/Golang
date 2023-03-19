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

// array_name := [length]Type{item1, item2, item3,...itemN}

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

	// 5. If an array does not initialize explicitly, then the default value of this array is 0.

	var givenValues [2]int
	fmt.Println("Elements of array are:", givenValues)

	// 6. In an array, you can find the length of the array using len() method

	arr1 := [5]int{3, 7, 1, 9, 10}

	// In an array, if ellipsis ‘‘…’’ become visible at the place of length, then the length of the array is determined by the initialized elements.
	arr2 := [...]int{1, 7, 3, 5, 8, 9, 24, 11}

	fmt.Println("Length of the array-01 is:", len(arr1))
	fmt.Println("Length of the array-02 is:", len(arr2))

	// 7. In an array, you are allowed to iterate over the range of the elements of the array.

	thisArray := [...]int{11, 24, 9, 17, 23, 25, 30, 44, 21, 20, 13, 19}

	for i := 0; i < len(thisArray); i++ {

		fmt.Println("Elements of array are:", thisArray[i])
		//OR
		//fmt.Printf("Elements of array are %d\n:", thisArray[i])
	}

	// 8. In Go language, an array is of value type not of reference type. So when the array is assigned to a new variable, then the
	// changes made in the new variable do not affect the original array

	oldArray := [...]string{"This", "arrays", "with", "Golang"}
	fmt.Println("Original array before:", oldArray)

	newArray := oldArray
	fmt.Println("New array before:", newArray)

	// Change the value at index 0 to Learning
	newArray[0] = "Learning"

	fmt.Println("New array after:", newArray)
	fmt.Println("Original array after:", oldArray)

	// 9. In an array, if the element type of the array is comparable, then the array type is also comparable. So we can directly compare
	// two arrays using == operator

	array1 := [3]int{9, 7, 6}
	array2 := [...]int{9, 7, 6}
	array3 := [3]int{9, 5, 3}

	// Comparing arrays using == operator
	fmt.Println(array1 == array2)
	fmt.Println(array2 == array3)
	fmt.Println(array1 == array3)

	// This will give and error because the
	// type of arr1 and arr4 is a mismatch
	// array4:= [4]int{9,7,6}
	// fmt.Println(array1==array4)

}
