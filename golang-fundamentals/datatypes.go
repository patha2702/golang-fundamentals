package main

import "fmt"

// Data types:
// These are the types of data that can be stored and used within the program
// Basic types:
//   1. Integer: These are whole numbers like 2, -8, 12, etc
//     There are two types of integers in go:
//       a. Signed int:
//	   Based on the size we have,
//	   - int, int8, int16, int32, int64
//	 b. Unsigned int:
//	   Based on the size we have
//	   - uint, uint8, uint16, uint32, uint64, uintptr
//   2. Float: These are decimal point numbers like 2.5, 23.45, etc
//	 Types: float32, float64
//   3. Complex numbers: For eg, 12 + 14i
//	 Types: complex64, complex128
//   4. string: These are text type and are enclosed by ". For eg, "rajendra"
//   5. bool: These are used to represent boolean values i.e. true, false
//   6. byte: alias for uint8
//   7. rune: alias for int32, represents unicode code point (one character in string)
//	  

// Note: Unless when performance and memory are your primary concerns, don't use types with size.
// Prefer using default ones like int, float64, etc

func main(){
	var totalCountOfStudents int = 60
	var billAmount uint = 10000
	var processingFee float64 = 0.2 * float64(billAmount)
	var hasFinishedPaying = true
	var statusMessage = "Transaction successful"
	fmt.Printf("%d %v %f %t %s", totalCountOfStudents, billAmount, processingFee, hasFinishedPaying, statusMessage)
}
