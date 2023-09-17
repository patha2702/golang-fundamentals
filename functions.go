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
}
