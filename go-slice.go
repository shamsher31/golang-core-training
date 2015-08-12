package main

import "fmt"

/**
 * Slice is a segment(part) of array
 * unlike array we can change length of slice
 */
func main() {

	/**
	 * make function is used to creat slices
	 * 1st param is type and 2nd param is length
	 */
	slice1 := make([]float64, 5)
	fmt.Println(slice1)

	/**
	 * you can also use arr[low,high] syntax
	 * arr[1:] = arr[1:len(arr)]
	 * arr[:5] = arr[0:5]
	 * arr[:] = arr[0:len(arr)]
	 * 'append' is like push at the end of the array
	 */
	slice2 := []float64{1, 2, 3}
	slice3 := append(slice2, 4, 5)
	slice4 := slice3[1:4]

	fmt.Println(slice3)
	fmt.Println(slice4)

	/**
	 * 'copy' will add first 2 elements of
	 * 1st slice into 2nd slice
	 */

	slice5 := []int{6, 7, 8, 9, 10}
	slice6 := make([]int, 2)
	copy(slice6, slice5)

	fmt.Println(slice5, slice6)

}
