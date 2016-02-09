package main

import "fmt"

// Function starts with keyword 'func'
// name of the function 'average'
// function parameter with following form
// name type, name type
// and return type before opening brackets

func average(axis []float64) float64 {
	total := 0.0

	for _, val := range axis {
		total += val
	}

	return total / float64(len(axis))

}

// Function can also return multiple values
// if you want to return multiple values
// you need to specify return type seprated
// by comma and return with multiple values

func f() (int, int) {
	return 10, 30
}

func main() {

	axis := []float64{98, 76, 47, 63, 28}
	fmt.Println(average(axis))

	x, y := f()
	fmt.Println(x, y)

	// This is an example of closure
	// if the params are of same type
	// you can write datatype only once
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(add(30, 20))

	// Function with multiple return value
	// writing func like this is called Closure
	// bcoz they are accessing x and y which
	// are local to func main and are not passed
	// as param to function

	sumAndProduct := func() (int, int) {
		return x + y, x * y
	}

	fmt.Println(sumAndProduct())

	// 'defer' moves function execution
	// at the end of last function call
	// 'recover' stops panic and return
	// value passed from panic

	defer func() {
		str := recover()
		fmt.Println(str)

	}()

	// it genrates run-time error
	panic("This is panic to indicate programmer error")

}
