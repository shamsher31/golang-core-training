// Ref: http://blog.golang.org/slices

package main

import (
	"bytes"
	"fmt"
)

// define path as a byte slice
type path []byte

func main() {

	var buffer [256]byte

	// Add 10 to last element of array
	buffer[255] = 10

	fmt.Println("Array buffer :", buffer)

	// create slice from array [start:end]
	slice := buffer[100:150]

	AddOneToEachElement(slice)
	fmt.Println("Part of the array : ", slice)

	SubstractOneFromLength(slice)
	fmt.Println("After substraction : ", slice)

	pathName := path("/use/local/bin")

	fmt.Println("Path : ", string(pathName))

	pathName.TruncateAtFinalSlash()
	fmt.Println("Modified path : ", string(pathName))

	var iBuffer [10]int

	slice4 := iBuffer[0:0]

	for i := 0; i < 11; i++ {
		slice4 = Extend(slice4, i)
		fmt.Println(slice4)
	}

}

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i] += 1
	}
}

func SubstractOneFromLength(slice []byte) []byte {
	// remove last elem of the slice
	slice = slice[0 : len(slice)-1]
	return slice
}

func (p *path) TruncateAtFinalSlash() {

	// Verify if the byte has a / in the end
	// and return its index, if not exist return -1
	i := bytes.LastIndex(*p, []byte("/"))
	fmt.Println("index", i)

	if i > 0 {
		// remmove value from last index
		*p = (*p)[0:i]
	}

}

func Extend(slice []int, elem int) []int {
	n := len(slice)
	slice = slice[0 : n+1]
	slice[n] = elem
	return slice
}
