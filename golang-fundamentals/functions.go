package main

import "fmt"


// Functions:
//   - A function is a block of code, that takes some input, does some processing and returns the result.
//   - It helps making the code reusable, i.e. once you define a function, you can it as many times as required
//     instead the writing the whole block of code again.

// Syntax:
//  func funName(paraName paraType) returnType {
//  	// Body
//      return result
//  }

// where,
// func funName(paraName paraType) returnType => function signature (it tells the caller about the function info
// like what arguments it takes, what it returns, it simply tells how one can use the function.)

func add(num1 int, num2 int) int {
        result := num1 + num2
        return result
}


// Multiple Parameters:
// When multiple arguments are of same type, the type only needs to be declared once after the last one, 
// assuming they are in order.

func greet(firstName, lastName string) string {
	greetMessage := fmt.Sprintf("Hello, %s %s, welcome to the platform, let's start the journey!!!", firstName, lastName)
	return greetMessage
}


// Pass by Value:
// Variables in Go are passed by value, that means when you pass an argument to a function, it receives a copy of the 
// passed data, and operates on that copied data.

func increment(num1 int) {
	num1 = num1 + 1
}


// Ignoring Return Values:
// A function can return multiple values, and some of them can be of no use for us in the program.
// Go doesn't allow you to have unused variables in the program.
// Storing the value in a variable, which is not likely to be used can result in an error like (variable declared but not used error).
// To deal with this, we can simply ignore the values which are not of use, by using '_'.

func getCoordinates() (int, int) {
	x := 2
	y := 1
	return x, y
}


// Named Return Values:
// Go's return vaules can be named. If done so, they are defined at the top of the function.
// By naming them they are initialised to zero value of the type.
// The return statement without any arguments, implicitly returns the named return values.
// This is called as naked return.
// They should be used only in short functions, otherwise they can harm readability of the program.

func subtract(num1, num2 int) (res int) {
	// res is initialised to zero value, here 0 since type is int.
	res = num1 - num2
	return // automatically returns res
}


// Explicit Returns:
// We can also explicitly return values with the named return values.
// Explicit Returns overrides the implicit return values.

func call() (message string) {
	message = "Hello, tring, tring"
	return message 
}


// Guard clause:
// Early return from the function when a given condition is met.
// They are used to check and handle exceptional cases upfront before executing the main logic.
// Helps to avoid deeply nested if-else loops.


func main(){
	num1 := 10
	num2 := 21
	result := add(num1, num2)
	fmt.Println(result)
	firstName := "Rajendra"
	lastName := "Patha"
	greeting := greet(firstName, lastName)
	fmt.Println(greeting)
	increment(num1)
	fmt.Println(num1) // num1 doesn't get incremented because it was passed by value
	
	// Ignoring Return Values
	x, _ := getCoordinates() // ignoring the y coordinate.
	fmt.Println(x)

	// Named return values
	difference := subtract(num1, num2)
	fmt.Println(difference)

	// Explicit Returns
	message := call()
	fmt.Println(message)
}
