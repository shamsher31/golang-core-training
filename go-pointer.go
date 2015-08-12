package main

import "fmt"

/**
 * Pointer points to memory location
 * '*' is pointer in go, it is also
 * used to refrence value stored in poniter
 * '&' is used to refrence pointer value
 */
func zipcode(p *int) {
	*p = 401203
}

func main() {

	pinCode := 411001

	zipcode(&pinCode)

	fmt.Println(pinCode)

}
