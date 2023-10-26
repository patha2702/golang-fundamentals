package main

import "fmt"

// Structs:
// - A struct is a composite data type that allows you to define your own custom types, 
//   by combining other types in a single entity.
// - It is a collection data type. That means it contains a collection of other data types.
// - It is similar to object literal in javascript and dictionary in python, and is used to 
//   encapsulate related data together.

// Syntax:
// type structName struct{
//    properties
// }

// Defining a car struct
type car struct {
	brand string
	model string
	cost float64
	color string
}


// Nested structs:
// In go we can have struct type inside another struct, this is called nested structs.
// For example:

type person struct{
	name string
	age int
	birthDate string
	location address // nested struct address
}

type address struct{
	city string
	state string
	country string
}


// Anonymous structs:
// - These are just like normal structs but are defined without a name, therefore they cannot be 
//   referenced or used else where in the code.
// - If a struct is being used only once in the program, then you can use/ create anonymous struct.


// Embedded structs:
// - Embedded structs provide a way for data only type of inheritance.
// - It allows us to share fields between struct definitions.


// Difference between nested and embedded structs:
// - Embedded structs fields are accessed at top level unlike nested struct
// - Embedded structs fields can be accessed like normal struct.

type sport struct {
	name string
	sportType string
	sportInfo // Embedded struct
}

type sportInfo struct {
	noOfPlayers int
	noOfReferees int
}


// Struct Methods:
// While Go is not objected oriented, it supports methods which can be defined on structs.
// Methods are just functions that have a receiver.
// Receiver => Special parameter that is defined before the name of function.

type rect struct {
	length  int
	breadth int
}

// (r rect) => receiver of area.
func (r rect) area() int {
	area := r.length * r.breadth
	return area
}

func main(){
	// Creating an instance of car struct
	benz := car{
		brand: "Benz",
		model: "S12",
		cost: 30.20,
		color: "Red",
	}

	// Accessing the fields of the struct using dot (.) operator
	fmt.Printf("Car brand: %v\n", benz.brand)
	fmt.Printf("Car model: %v\n", benz.model)
	fmt.Printf("Car cost: %v\n", benz.cost)
	fmt.Printf("Car color: %v", benz.color)

	// Creating an instance of Nested struct
	raj := person{
		name: "Raj",
		age: 1,
		birthDate: "12/12/2022",
		location: address{
			city: "Pune",
			state: "Maharashtra",
			country: "Bharat",
		},
	}
	
	fmt.Println(raj)

	// Alternative way of creating an instance.
	// After instanciating struct this way, all property values will be initialised to 
	// zero values of their types.

	raju := person{}
	fmt.Println(raju)
	raju.name = "Raju"
	raju.age = 12
	raju.birthDate = "20/01/2011"
	raju.location.city = "Hyderabad"
	raju.location.state = "Telangana"
	raju.location.country = "Bharat"
	
	fmt.Println(raju)


	// Creating anonymous struct
	// Immediately instanciating a struct after defining it.
	bag := struct{
		color string
		volume float64
		cost int
	}{
		color: "blue",
		volume: 3.5,
		cost: 750,
	}

	fmt.Println(bag)


	// Embedded structs
	cricket := sport{
		name: "Cricket",
		sportType: "Outdoor",
		// In embedded structs key is same as embedded struct name.
		sportInfo: sportInfo{
			noOfPlayers: 11,
			noOfReferees: 3,
		},
	}
	fmt.Println(cricket)
	

	// Struct Methods:

	myRect := rect{
		length:  25,
		breadth: 10,
	}

	myRectArea := myRect.area()
	fmt.Println(myRectArea)
}
