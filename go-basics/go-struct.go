package main

import (
	"fmt"
	"math"
)

/**
 * 'type' defines new type
 * 'Circle' is name of struct
 */
type Circle struct {
	x, y, r float64
}

/**
 * area() is name of function
 * float64 is a return type
 * (c *Circle) is called reciver
 * Reciver is like paramter(name type)
 * By using this method we can call
 * func using dot operator
 */
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {

	/**
	 * Initilize circle struct with default values
	 */
	c := Circle{x: 0, y: 0, r: 2}

	fmt.Println(c.area())

}
