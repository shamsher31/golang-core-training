package main

import (
	"fmt"
	"sort"
)

func min(n []int) int {
	sort.Sort(sort.IntSlice(n))
	return n[0]
}

func max(n []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	return n[0]
}

func main() {
	num := []int{
		94, 67, 86, 58, 89, 43, 20, 99,
	}

	fmt.Println(min(num))
	fmt.Println(max(num))

}
