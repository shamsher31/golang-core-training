package main

import "fmt"

/**
 * This variable will be accesible to other function in the same package
 */
var country = "India"

/**
 * The go compiler will not allow you to create
 * varible that you never use and will through
 * 'variable declared and not use' error
 */
func main() {

	/**
	 * Variable declaration has 3 parts
	 * keyword 'var' is used to declare variable
	 * Name of the variable
	 * type of variable
	 */
	var name string
	name = "Shamsher Ansari"
	fmt.Println("Full name : ", name)

	/**
	 * const is used to declare constants
	 * whos value can not be changed throught the application
	 */
	const dob = "1989/03/31"
	fmt.Println("Date of Birth : ", dob)

	/**
	 * += is used on concatenation
	 */
	var city string
	city = "Mumbai"
	city += " Pune"
	fmt.Println("City : ", city)

	/**
	 * You can use ':' which automatically gues type of value
	 */

	zipcode := 411001
	fmt.Println("Zipcode : ", zipcode)

	/**
	 * () are used to call function defined in the same package
	 */
	place()

}

func place() {
	fmt.Println("Country : ", country)
}
