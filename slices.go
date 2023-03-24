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

	// 7.

	// =-=-=-=-=-=-=-=-=-=-Using make() function:=-=-=-=-=-=---=--=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-

	//You can also create a slice using the make() function which is provided by the go library. This function takes
	// three parameters, i.e, type, length, and capacity. Here, capacity value is optional. It assigns an underlying array with a size that is
	// equal to the given capacity and returns a slice which refers to the underlying array. Generally, make() function is used to create an empty
	// slice. Here, empty slices are those slices that contain an empty array reference.

	// Syntax:

	// func make([]T, len, cap) []T

	// Creating an array of size 7 and slice this array  till 4 and return the reference of the slice Using make function

	slice_1 := make([]int, 5, 7)
	fmt.Printf("slice_1= %v\n, length= %d\n, capacity= %d\n", slice_1, len(slice_1), cap(slice_1))

	// Creating another array of size 7 and return the reference of the slice Using make function

	var slice_2 = make([]int, 7)
	fmt.Printf("Slice 2 = %v, \nlength = %d, \ncapacity = %d\n", slice_2, len(slice_2), cap(slice_2))

	// 8.

	// How to iterate over a slice?

	// (A) Using for loop:

	myslice := []string{"This", "is", "slice", "tutorial", "of", "Go", "language"}

	// Iterate using for loop
	for i := 0; i < len(myslice); i++ {

		fmt.Println(myslice[i])
	}

	// (B) Using range in for loop:
	//Using range in the for loop, you can get the index and the element value:

	thisSlice := []string{"This", "is", "slice", "tutorial", "of", "Go", "language"}

	for index, element := range thisSlice {

		fmt.Printf("Index = %d and element = %s\n", index, element)
	}

	// 9.

	// Using a blank identifier in for loop:

	// In the range for loop, if you don’t want to get the index value of the elements then you can use
	// blank space(_) in place of index variable :

	givenSlice := []string{"This", "is", "slice", "tutorial", "of", "Go", "language"}

	for _, elem := range givenSlice {

		fmt.Printf("element = %s\n", elem)
	}

	// 10.

	// =-=--=-=-=-=-=-=--=--=-=-=--=-=-=- [] IMPORTANT POINTS [] =-=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=--=

	// [] Zero value slice:

	// In Go language, you are allowed to create a nil slice that does not contain any element in it. So the capacity and the length of this slice
	// is 0. Nil slice does not contain an array reference

	var slicing []string
	fmt.Printf("Length = %d\n", len(slicing))
	fmt.Printf("Capacity = %d ", cap(slicing))

	// [] Modifying Slice:
	// As we already know that slice is a reference type it can refer an underlying array. So if we change some elements in the slice, then the
	// changes should also take place in the referenced array. Or in other words, if you made any changes in the slice, then it will also reflect
	//  in the array.

	integerArray := [8]int{45, 21, 27, 11, 19, 41, 38, 76}

	slicedIntegerArray := integerArray[0:6]

	fmt.Println("\nInteger Array:", integerArray)
	fmt.Println("Sliced Integer Array:", slicedIntegerArray)

	slicedIntegerArray[0] = 171
	slicedIntegerArray[1] = 109
	slicedIntegerArray[2] = 348
	slicedIntegerArray[3] = 225

	fmt.Println("Integer Array:", integerArray)
	fmt.Println("Sliced Integer Array:", slicedIntegerArray)

	// [] Comparison of Slice:
	// In Slice, you can only use == operator to check the given slice is nill or not. If you try to compare two slices with the help of ==
	// operator then it will give you an error Comparison of Slice: In Slice, you can only use == operator to check the given slice is nill
	//or not. If you try to compare two slices with the help of == operator then it will give you an error Comparison of Slice: In Slice, you
	// can only use == operator to check the given slice is nill or not. If you try to compare two slices with the help of == operator then it
	// will give you an error

	s1 := []int{5, 6, 8, 11}
	var s2 []int
	s3 := []int{13, 7, 1, 5}

	//fmt.Println(s1 == s3) --> Error: invalid operation: s1 == s3 (slice can only be compared to nil)

	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	fmt.Println(s3 == nil)

	// [] Note: If you want to compare two slices, then use range for loop to match each element or you can use DeepEqual function.[]

	// []Multi-Dimensional Slice:
	// Multi-dimensional slice are just like the multidimensional array, except that slice does not contain the size.

	// [] M1:
	multiSlice_01 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Print("Multi-Dimensional Slice-01:", multiSlice_01)

	// [] M2:

	multiSlice_02 := [][]string{[]string{"learn", "golang"}, []string{"learn slice"}, []string{"slice", "array"}}

	fmt.Println("\nMulti-Dimensional Slice-02:", multiSlice_02)

	// [] Sorting of Slice:

	// In Go language, you are allowed to sort the elements present in the slice. The standard library of Go language provides the sort package
	// which contains different types of sorting methods for sorting the slice of ints, float64s, and strings. These functions always sort the
	// elements available is slice in ascending order.

	sortSlice_01 := []string{"Python", "Java", "C#", "Go", "Swift"}
	sortSlice_02 := []int{38, 41, 22, 15, 17, 40, 33, 91}

	fmt.Println("Before sorting slices are :")

	fmt.Println("\nsortSlice_01", sortSlice_01)
	fmt.Println("sortSlice_02", sortSlice_02)

	sort.Strings(sortSlice_01)
	sort.Ints(sortSlice_02)

	fmt.Println("After sorting slices are :")

	fmt.Println("\nsortSlice_01", sortSlice_01)
	fmt.Println("sortSlice_02", sortSlice_02)
}
