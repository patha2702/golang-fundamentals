package main

import (
	"fmt"
)

// Arrays:
// - Array is an ordered list/ collection of similar items
// - These are fixed-size groups of variables of same type.


// Slices:
// - Dynamically-sized, flexible view onto the array. Similar to dynamic arrays in c++.
// - These are just reference to the part of specified array's elements
//   that means changing the elements in the slices is also reflected in original array.
// - slice1 := arr[low:high], low => inclusive, high => exclusive
// - nil is the zero value of slice


// make([]T, len, cap) []T:
// - We can create a new slice using make() function
// - Slices created with make() will be filled with zero value of the type.


// Variadic functions:
// - These are used to take variable arguments like *args in python
// - It uses ... => ellipsis notation. Eg ...int

func sum(nums ...int) int{
	// nums ...int => slice
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}


// Spread operator:
// - The spread operator allows us to pass a slice into a variadic function while calling.
// - eg. nums...


// append(slice []T, elements ...T) []Type => To add elements to slice


func main() {
	// Creating array:
	var myArr [10]int // values are initialized to zero value of the type
	fmt.Printf("%T", myArr)

	myArr2 := [5]int{1,3,5,7,6}
	fmt.Println(myArr2)

	// Creating slices:
	mySlice := myArr2[2:5]
	fmt.Println(mySlice)

	var mySlice2 []int 
	mySlice2 = myArr2[:]
	mySlice2[1] = 100
	
	// Creating slice using make() function:
	mySlice3 := make([]int, 5, 10)
	fmt.Println(mySlice3)

	// Variadic functions (passing values using slice spread operator)
	nums := []int{1,2,3,4,5}
	sum := sum(nums...)
	fmt.Println(sum)

	// Appending elements to slice
	nums = append(nums, 20,30,40)
	fmt.Println(nums)
}
