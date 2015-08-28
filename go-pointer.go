package main

import "fmt"

// Pointer points to memory location
// '*' is pointer in go, it is also
// used to refrence value stored in poniter
// '&' is used to refrence pointer value

func zipcode(p *int) {
	*p = 401203
}

// if you want to change the original
// value of param then use pointer
func add(a *int) int {
	*a = *a + 1
	return *a
}

func main() {

	pinCode := 411001
	zipcode(&pinCode)
	fmt.Println(pinCode)

	x := 5
	fmt.Println("old value of x:", x)

	// passing x using &, pass the address of x
	// not a copy of x, so any changes to param in add
	// func will change the original value
	x1 := add(&x)
	fmt.Println("new value of x:", x1)

}
