// compare the age of two people, then return the older person and differences of age
package main

import "fmt"

// person is a new type of struct
type person struct {
	name string
	age  int
}

// struct are pass by value
// you can pass person as prams
// and specify as return type
func Elder(p1, p2 person) (person, int) {

	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age

}

func main() {

	// create and asign default values
	sam := person{name: "Sam", age: 25}
	kim := person{"Kim", 35}
	tom := person{"Tom", 30}

	// Call function and pass struct
	fmt.Println(Elder(sam, kim))
	fmt.Println(Elder(sam, tom))
	fmt.Println(Elder(kim, tom))

}
