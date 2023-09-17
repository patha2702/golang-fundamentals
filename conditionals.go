package main

import "fmt"

// Conditionals:
// if - else:
// These are used to execute statements if certain condition is met.
// If the expression in 'if' construct evaluates to true, then 
// statements inside if block are executed.
// If not evaluates to true, then it checks expression in 'else if' construct, 
// if it is evaluated to true, statements inside else if block are executed.
// If no condition is evaluated to true, then else block is executed.


// Syntax:
// if cond/expression {
//   statements
// }else if{
//   statements
// } else {
//   statements
// }


// Initial statement of an if block:
//   -When to use: If you are declaring a variable, that is only used to check certain condition,
//      and not used else where in program, then prefer using initial statement of if block.
//   -To keep code precise and limit the scope.

// Syntax:
// if initial statement; condition {
//   statements
// }


func main(){
	percent := 88

	if percent >= 75 {
		fmt.Println("You have passed with distinction")
	} else {
		fmt.Println("Sorry, you have missed the distinction, better luck next time.")
	}

	// Initial statement of if block
	password := "Raj1234"

	if length := len(password); length < 5 {
		fmt.Println("Password is too small.")
	} else {
		fmt.Println("Password is valid.")
	}
}


