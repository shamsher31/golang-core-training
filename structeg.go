package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	fmt.Println(Person{"Sam", 20})
	// {Sam 20}

	fmt.Println(Person{name: "Shamsher", age: 23})
	// {Shamsher 23}

	fmt.Println(Person{name: "Shams"})
	// {Shams 0}

	fmt.Println(&Person{name: "Sham"})
	// &{Sham 0}

	p := Person{name: "Pranay", age: 23}
	fmt.Println(p.name)
	// Pranay

	pp := &p
	fmt.Println(pp.age)
	// 23
}
