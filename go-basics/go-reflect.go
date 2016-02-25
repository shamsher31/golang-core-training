package main

// package reflect provide runtime reflection
// it allows us to identify data type of value
import (
	"fmt"
	"reflect"
)

type Contact struct {
	Name string `json:"name"`
	Id   int    `json:id`
}

func main() {
	x := 7.5
	y := float32(5.5)
	name := "Shamsher"
	age := 27

	fmt.Printf("var x %v = %v \n", reflect.TypeOf(x), x)
	fmt.Printf("var y %v = %v \n", reflect.TypeOf(y), y)

	// ValueOf returns reflect.Value which hold the name
	// not the value of name
	person := reflect.ValueOf(name)

	personAge := reflect.ValueOf(age)

	// type of person is reflect.Value which holds name
	fmt.Printf("valueOf person is : %v \n", reflect.TypeOf(person))
	// we have to convert person to string to get its value
	fmt.Println("Content of the name is : ", person.String())

	fmt.Printf("valueOf age is : %v \n", reflect.TypeOf(personAge))

	// reflect.Value has several methods
	// reflect.Value.Bool(), reflect.Value.Int(), reflect.Value.float()
	fmt.Println("Content of the age is :", personAge.Int())

	// set the default value
	employee := Contact{"Shams", 3283283}

	// It returns struct type as main.Contact
	employeeType := reflect.TypeOf(employee)
	fmt.Println(employeeType)

	if fields, ok := employeeType.FieldByName("Name"); ok {
		fmt.Printf("%q %q %q \n", fields.Name, fields.Type, fields.Tag)
	}
}
