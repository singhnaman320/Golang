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
	fmt.Println("Slices:", sort.IntsAreSorted(slices))

	// 3.

}
