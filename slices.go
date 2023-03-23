// Slices in Go are a flexible and efficient way to represent arrays, and they are often used in place of arrays because of their dynamic size
// and added features. A slice is a reference to a portion of an array. It’s a data structure that describes a portion of an array by specifying
// the starting index and the length of the portion

// The first index position in a slice is always 0 and the last one will be (length of slice – 1)

package main

import (
	"fmt"
	"sort"
)

func main() {

	// 1.
	array := [5]int{11, 22, 33, 44, 55}
	slice := array[1:4]

	fmt.Println("Array:", array)
	fmt.Println("Slice:", slice)

	// 2.

	// Slices are dynamic, which means that their size can change as you add or remove elements. Go provides several built-in functions that allow
	// you to modify slices, such as append, copy, and delete.

	// Here, the function append is used to add the elements 15, 22, 26 to the slice slice. The result is a new slice that contains the elements
	// 14, 17, 11, 16, 15, 22, 26

	slices := []int{14, 17, 11, 16}
	slices = append(slices, 15, 22, 26)

	fmt.Println("Slices:", slices)
	fmt.Println("Slices:", sort.IntsAreSorted(slices)) //false

	// 3.
	// Declaration of Slice
	// A slice is declared just like an array, but it doesn’t contain the size of the slice. So it can grow or shrink according to the requirement.

	// Syntax:

	// []T
	// or
	// []T{}
	// or
	// []T{value1, value2, value3, ...value n}
	// Here, T is the type of the elements.

	//Example:
	// var my_slice[]int

	// [] Components of Slice
	// A slice contains three components:

	// Pointer: The pointer is used to points to the first element of the array that is accessible through the slice. Here, it is not necessary that
	// the pointed element is the first element of the array.
	// Length: The length is the total number of elements present in the array.
	// Capacity: The capacity represents the maximum size upto which it can expand.

	letters := []string{"start", "learning", "slice", "with", "golang"}

	fmt.Println("Array is:", letters)

	letter_slice := letters[1:5]

	fmt.Println("slice array is", letter_slice)

	fmt.Printf("Length of slice: %d\n", len(letter_slice)) //4

	fmt.Printf("Capacity of slice: %d\n", cap(letter_slice)) //4

	//Explanation: In the above example, we create a slice from the given array. Here the pointer of the slice pointed to index 1 because the
	// lower bound of the slice is set to one so it starts accessing elements from index 1. The length of the slice is 4, which means the total
	// number of elements present in the slice is 5 and the capacity of the slice 5 means it can store a maximum of 6 elements in it.

	// 4.

	//How to create and initialize a Slice?

	// In Go language, a slice can be created and initialized using the following ways:

	// [] Using slice literal:

	// You can create a slice using the slice literal. The creation of slice literal is just like an array literal, but with one difference you
	// are not allowed to specify the size of the slice in the square braces[]. As shown in the below example, the right-hand side of this
	// expression is the slice literal.
	// var my_slice_1 = []string{"Geeks", "for", "Geeks"}
	// Note: Always remember when you create a slice using a string literal, then it first creates an array and after that return a slice
	// reference to it.

	var firstSlice = []string{"create", "and", "initialize", "a", "Slice"}

	fmt.Println(firstSlice)

	var secSlice = []int{22, 16, 18, 48, 76}

	fmt.Println(secSlice)

	// 5.

	// Using an Array: As we already know that the slice is the reference of the array so you can create a slice from the given array.

	// Syntax:

	// array_name[low:high]

	// Note: The default value of the lower bound is 0 and the default value of the upper bound is the total number of the elements present in
	// the given array.

	words := [5]string{"create", "and", "initialize", "Golang", "Slice"}

	sliceOne := words[1:2]
	sliceTwo := words[0:]
	sliceThree := words[:4]
	sliceFour := words[:]

	fmt.Println("First slice", sliceOne)
	fmt.Println("Second slice", sliceTwo)
	fmt.Println("Third slice", sliceThree)
	fmt.Println("Four slice", sliceFour)

	// 6.

	// Using already Existing Slice: It is also be allowed to create a slice from the given slice. For creating a slice from the given slice
	// first you need to specify the lower and upper bound, which means slice can take elements from the given slice starting from the lower bound
	// to the upper bound. It does not include the elements above from the upper bound.

	// Syntax:

	// slice_name[low:high]

	// This syntax will return a new slice.

	// Note: The default value of the lower bound is 0 and the default value of the upper bound is the total number of the elements present in
	// the given slice.

	// Creating s slice
	original_slice := []int{90, 60, 40, 50, 34, 49, 30}

	slice_One := original_slice[1:2]
	slice_Two := original_slice[0:]
	slice_Three := original_slice[2:5]
	slice_Four := original_slice[:]

	fmt.Println("Original slice", original_slice)

	fmt.Println("First slice", slice_One)
	fmt.Println("Second slice", slice_Two)
	fmt.Println("Third slice", slice_Three)
	fmt.Println("Four slice", slice_Four)
}
