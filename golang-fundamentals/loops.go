package main

import (
	"fmt"
	"math"
)


// Loops:
// Go has only one loop i.e. for loop.
// Syntax
// for init; cond; post {
//  // statements
// }
// Note: No parentheses in for loop signature.

// init => executes before the first iteration, and can create variables within the scope of loop.
// cond => evaluated before every interation, if condition isn't passed loop breaks.
// post => executes at the end of every iteration.


// Omitting conditions from a for loop:
// - We can omit sections of a for loop in GO.


// Go for loop is other languages while loop, because we can omit sections of for loop


// continue keyword:
// - continue keyword stops the current iteration and continues to the next iteration


// break keyword:
// - break keyword breaks the current iteration and exits the loop.


func isPrime(num float64) bool {
	// Even number cannot be prime
	if int(num) % 2 == 0 || num == 1.0{
		return false
	}
	
	// Check any number in range(3, sqrt(numer)) that can divide num
	for i := 3; i <= int(math.Sqrt(num)); i++ {
		if int(num) % i == 0 {
			return false
		}
	}
	
	// num is not even and is not divisible by any other number, hence prime
	return true
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i+1)
	}

	// Omitting conditions from a for loop:
	// - We can omit sections of a for loop in GO.
	
	// Omitting condition section to create infinite loop
	// for i := 0; ; i++ {
	//	fmt.Println(i)
	//}
	
	// Go for loop as while loop, by omitting init and post section
	count := 0
	for count < 100{
		fmt.Println("Count is: ", count)
		count++
	}

	// Skip multiples of 3 in range 1 - 100
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 {
			continue
		}
		fmt.Println(i)
	}	

	// Skip numbers that are not prime
	
	for i := 1; i <= 100; i++ {
		if ! isPrime(float64(i)) {
			continue
		}

		fmt.Println(i)
	}
}
