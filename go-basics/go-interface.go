package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!!"
}

type Student interface {
	getFullName() string
	getClass() string
	getRollNo() int
}

type student struct {
	FullName string
	Class    string
	RollNo   int
}

func (s *student) getFullName() string {
	return s.FullName
}

func (s *student) getClass() string {
	return s.Class
}

func (s *student) getRollNo() int {
	return s.RollNo
}

func main() {
	animals := []Animal{Dog{}, Cat{}}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

	std := student{"Shamsher Ansari", "BE", 24}
	fmt.Println("<<<< Student Details >>>>")
	fmt.Println("Full Name : ", std.getFullName())
	fmt.Println("Class : ", std.getClass())
	fmt.Println("Roll No. : ", std.getRollNo())

}
