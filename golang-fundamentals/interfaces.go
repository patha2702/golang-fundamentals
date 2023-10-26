package main

import (
	"fmt"
	"math"
)


// Interfaces:
// An interface type is defined as a set of method signatures.
// A value of interface type can hold any value that implements those methods.
// - A type implements an interface if it has all the methods of the given interface defined on it.

// - The following example means that for a type to be shape, it must have methods area() and perimeter()
//   defined on it.

// Interfaces enable:
// - Polymorphism: You can write code that works with different types as long as they satisfy the interface. 
//   This makes the code more flexible and allows you to build generic functions and libraries.

// - Abstraction: By defining interfaces, you can focus on what methods a type should implement 
//   without being concerned about the specific implementation details.


// Interface implementation:
// Interfaces in Go are implicitly implemented if a type has all the required interface methods defined on it.
// A type can implement any number of interfaces in Go.

type shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	area := math.Pi * c.radius * c.radius
	return area
}

func (c circle) perimeter() float64 {
	perimeter := 2 * math.Pi * c.radius
	return perimeter
}

func getArea(s shape){
	fmt.Printf(" Area is: %v", s.area())
}


// Naming the interface arguments
// - It makes the method signatures more readable, i.e. what arguments to pass, and what to expect as return value.

type copier interface {
	copy(sourceFile string, destinationFile string) (bytesCopied int)
}


// Type assertions:
// - It is used to check whether a value stored by the interface is of particular type or not.
// - t, ok := u.(T)
// - The above assertion checks that value of interface u is of type T or not.
// - It returns two values:
// -  1. original value t, if assertion is true, else t will be zero value of type T
//    2. boolean value reporting whether the assertion succeeded or not, stored in ok


// Type switches:
// - It makes easy to do several type assertions in a series.


func main() {
	// Here shape interface can store circle value, since circle implements shape interface.
	var s shape
	s = circle{
		radius: 2.5,
	}

	getArea(s)

	// Type assertion
	var i interface{} = "Hello"
	t, ok := i.(float64)
	fmt.Println(t, ok)

	// Type switches

	var num interface{} = true
	switch v := num.(type) {
		case string:
			fmt.Println("Interface contains a string value")
		case int:
			fmt.Println("Interface contains a integer value")
		case float64:
			fmt.Println("Interface contains a float value")
		default:
			fmt.Printf("Interface is of type %T", v)
	}
}
