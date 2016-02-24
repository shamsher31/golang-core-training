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

// it takes all var as *int( pointer to int)
// it modify and returns final product
func swapAndProduct1(x, y, product *int) {
	if *x > *y {
		*x, *y = *y, *x
	}

	*product = *x * *y
}

func main() {

	pinCode := 411001
	fmt.Println("Before : ", pinCode)
	zipcode(&pinCode)
	fmt.Println("After : ", pinCode)

	x := 5
	fmt.Println("old value of x:", x)

	// passing x using &, pass the address of x
	// not a copy of x, so any changes to param in add
	// func will change the original value
	x1 := add(&x)
	fmt.Println("new value of x:", x1)

	p1 := 9
	p2 := 5
	product := 0

	fmt.Println("Before: ", p1, p2, product)

	swapAndProduct1(&p1, &p2, &product)

	fmt.Println("After : ", p1, p2, product)

}
