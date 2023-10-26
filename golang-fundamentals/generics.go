package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)


// Generics:
// - It is used to write DRY code.
// - For eg, before generics you would have to write 2 separate functions 
//   for splitting a slice of int, and slice of string.
// - Because we need to specify the type at compile time.
// - Using generics you can write a single function which can take any parameter type.


// Type parameters:
// - T => Type parameter in uppercase, is a placeholder which will be replaced by the type of
//      value passed at compile.

// - any => Keyword specifying, that it can be any type, such as string, int, float64


func split[T any](s []T) ([]T, []T) {
        mid := len(s)/2
        return s[:mid], s[mid:]
}


// Type constraints:
// - It constraints/ restricts the types that can be used in generics
// - Type constraints allow the compiler to enforce type safety and ensure that only compatible
//   types are used with a generic construct.
// - Type constraints are specified using interface keyword


func max[T constraints.Ordered](a,b T) T {
    	if a > b {
		return a
	}
	return b
}

// Here, comparable is a built-in interface, which requires the type to implement comparison operators.
// This means, T is restricted to types that are comparable. 
// Note: All types that are comparable implicitly implement comparable interface


func main() {
	// Type parameters
	mySlice := []int{1,2,3,4,5}
	mySlice2 := []string{"Maharashtra", "Kerala", "Mumbai", "Delhi"}
	firstHalf, secondHalf := split(mySlice)
	firstHalf1, secondHalf1 := split(mySlice2)
	fmt.Println(firstHalf, secondHalf)
	fmt.Println(firstHalf1, secondHalf1)

	// Type constrainsts
	num1 := 12
	num2 := 23
	res := max(num1, num2)
	fmt.Println(res)
}
