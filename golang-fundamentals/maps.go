package main

import (
	"fmt"
)


// Maps:
// - These are similar to dictionary in python.
// - It provides key-value mapping.
// - We can create map using 
//   a. literal
//   b. make() function
// - zero value of maps => nil
// - They are passed by reference.

// Operations on maps:
// 1. Insert an element:
//   m[key] = value

// 2. Get an element:
//   element = m[key]

// 3. Delete an element
//    delete(m, key)

// 4. Check if an element exiss
//    element, ok := m[key]
//    Note: if key in m then ok is true, else false and element value will be zero value of map's element.

// Accessing the key-value pair not present in the map, returns the zero value.


// Exercise
func getNameCount(names []string) map[string]map[string]int {
	nameCount := make(map[string]map[string]int)
	for _, name := range names {
		initial := string(name[0])
		// If initial not present, add initial along with name and count.
		if _, ok := nameCount[initial]; !ok {
			nameCount[initial] = map[string]int {
				name : 1,
			}
			continue
		}
		// Initial is there, but name is not there, then add the name and count.
		if _, ok := nameCount[initial][name]; !ok {
			nameCount[initial][name] = 1
		} else {
			// Same name is present already, then increment the count
			nameCount[initial][name] += 1
		}
	}

	return nameCount
}


func main() {
	// Maps using literal
	myMap1 := map[string]int{
		"Delhi": 21,
		"Maharashtra": 12,
		"Goa": 10,
	}
	fmt.Println(myMap1)

	// Maps using make()
	myMap2 := make(map[int]string)


	// Len function returns the total key-value pairs in map
	fmt.Println(len(myMap1))

	// Operations:
	// Insert:
	myMap2[10] = "ten"
	myMap2[9] = "nine"

	// Retrieval:
	element := myMap2[9]
	fmt.Println(element)

	// Delete:
	delete(myMap2, 9)

	// Check if present or not:
	el, ok := myMap2[10]
	el2, ok := myMap2[9]
	fmt.Println(el, ok)
	fmt.Println(el2, ok)

	// Iterating over the map
	for key := range myMap1 {
		fmt.Println(myMap1[key])
	}

	// Exercise 
	usernames :=  []string{"ram", "sam", "raj", "ram", "shyam", "shyam", "krishna"}
	result := getNameCount(usernames)
	fmt.Println(result)
}
