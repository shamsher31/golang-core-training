package main

import (
	"fmt"
)

/**
 * 'maps' is like associative array
 * Also called as Hash or Dictionary
 * map should be initilized using 'make' before use
 * maps are basically used as lookup table or dictionary
 */
func main() {

	/**
	 * Initilize map with key as a string
	 * and value as a string. you can also
	 * use map of key and val as a int
	 */
	days := make(map[string]string)
	days["Sun"] = "Sunday"
	days["Mon"] = "Monday"
	days["Tue"] = "Tuesday"
	days["Wed"] = "Wednesday"
	days["Thu"] = "Thursday"
	days["Fri"] = "Friday"
	days["Sat"] = "Saturday"

	// Print all values of map
	fmt.Println(days)

	// Print single value of map
	fmt.Println(days["Sun"])

}
