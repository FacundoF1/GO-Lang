package examples

import "fmt"

func TryArrays() {

	// Array
	var emptyArray [5]int
	fmt.Printf("An empty array: %v\n", emptyArray)

	// Assigning to an array
	var arrayWithValues [2]int
	arrayWithValues[0] = 10
	arrayWithValues[1] = 20
	fmt.Printf("An array with values: %v\n", arrayWithValues)

	// Multidimensional array
	var matrix [2][2]int
	matrix[0][0] = 1
	fmt.Printf("A matrix with values: %v\n", matrix)

	// Slices
	var slice []int
	slice = make([]int, 3)
	fmt.Printf("An empty slice: %v\n", slice)
	slice[0] = 0
	slice[1] = 1
	slice[2] = 2

	slice = append(slice, 3)
	fmt.Printf("A populated slice: %v\n", slice)

	// Range over a slice
	for index, value := range slice {
		fmt.Printf("Ranging over slice, at index %v with value %v\n", index, value)
	}
}
