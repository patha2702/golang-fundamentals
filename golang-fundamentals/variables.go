package main

import "fmt"

// A variable is a reserved memory space that can contain values, which can be used later when needed.
// It can be imagined as box with a label indicating what is inside the box.
// studentAge = 21, 
//   (label)           --------------   
//  studentAge ------->|     21	    |
// 		       |            |
//                     -------------

// Variable declaration:
// var varname datatype = value
// For eg, var myname string = "Raj"
// By default, when we declare the variable, zero equivalent of the type is assigned.

// To simplify the variable declaration, there is also an alternative: Short Assignment Operator
// varname := value
// For eg, username := "rajen123"
// Here, the type is determined by the value stored in it at compile time.
// Unlike in python, you cannot change the value to different type, once declaring 
// as string.
// For eg, username = 23 is not allowed.
// ':=' cannot be used at global scope, i.e. outside the function in the package.


func main(){
	var username string = "patha2702"
	password := "1234567890"
	
	// Multiple declarations on same line
	name, age := "raj", 21
	fmt.Println("Basic authentication " + username + ":" + password)
	fmt.Printf("Name is %s, age %d", name, age)
}
