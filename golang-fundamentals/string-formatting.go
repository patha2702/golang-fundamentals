package main

import "fmt"

// String Formatting:
// 1. fmt.Printf:
//   - It prints the formatted string to the standard output/ console.

// 2. fmt.Sprint:
//   - It returns the formatted string.

// String interpolation formatting verbs:
//   - %v : To interpolate value
//   - %s : for string
//   - %d : for integer
//   - %f : for float value
//   - %t : for boolean
//   - %T : for the type of the value

func main(){
	country := "India"
	states := 29
	growthRate := 12.5
	potentialToGrow := true
	fmt.Printf("%s has %d states with potential to grow %t, at %f growth rate.", country, states, potentialToGrow, growthRate)
	message := fmt.Sprintf("I am from %v, type: %T", country, country)
	fmt.Print(message)
}
