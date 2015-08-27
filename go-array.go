package main

import "fmt"

func main() {

	// Array in go must be of fixed length and same type
	// You can write array in one line or split it like following
	// Notice ',' at the end of array. This is required by go
	// Compiler so you can remove and comment elements from array
	x := [5]float32{
		32,
		43,
		54,
		76,
		21,
		// 34,
	}

	//Decalre variable to hold average with same data type as array above
	var total float32 = 0

	// You can calculate the length of array using 'len()'
	// range keyword avoids calculating len of array
	// '_' tells compiler that this variable will not use
	// in loop, if you write 'i' instead of '_'
	// go compiler will through 'value declared and not use' error
	for _, value := range x {
		total += value
	}

	// Typecaset array length bcoz len returns int type
	// and we need float32
	fmt.Println(total / float32(len(x)))

	// Create 2-D array of 2 elem each with 4 elem
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(easyArray)

	// You can also represent above array as
	hardwayArray := [2][4]int{[4]int{9, 10, 11, 12}, [4]int{12, 14, 15, 16}}
	fmt.Println(hardwayArray)

}
