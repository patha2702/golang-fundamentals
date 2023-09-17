package main

import "fmt"

// Constants:
// These are values that once assigned, cannot be changed later.
// Value is fixed.

// Syntax:
// const identifier type = value
// Note: You can't use := operator for constants.
// Constant can be assigned an expression, but the value should be known at compile time.

func main(){
	const monthsInAYear int = 12
	const monthsInDecade int = monthsInAYear * 10
	fmt.Println(monthsInDecade)
	fmt.Println(monthsInAYear)
}
