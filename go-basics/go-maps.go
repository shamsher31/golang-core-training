package main

import "fmt"

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
	 * Use better way of creating maps
	 */
	days := map[string]string{
		"Sun": "Sunday",
		"Mon": "Monday",
		"Tue": "Tuesday",
		"Wed": "Wednesday",
		"Thu": "Thursday",
		"Fri": "Friday",
		"Sat": "Saturday",
	}

	// Print all values of map
	fmt.Println(days)

	// Print single value of map
	fmt.Println(days["Sun"])

	/**
	 * Go provides better way to access map values
	 * 'name' is the lookup result
	 * 'ok' returns boolen and tells whether or not
	 * lookup was successful
	 */
	if name, ok := days["Fri"]; ok {
		fmt.Println(name, ok)
	}

}
