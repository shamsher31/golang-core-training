// Package Declaration
// Every Go programe must start with Package declaration
// Packages are Go's way of organizing and re-using code

package main

// Import keyword is used to include other packages
// fmt(format) package is used for formatting
// You can import packages in 3 way
// using '.' will omit package name so you can write Println instead of fmt.Println
// using 'aliase' so you can write f "fmt" and write f.Println
// using '_' like _ "fmt" so you can use only init() from package
// and you can not use methods from package

import (
	"fmt"
	"time"
)

// func is keyword for writing function
// Println(Print Line)
// No semi colon at the end of line

func main() {
	fmt.Println("My first awesome go programe")

	fmt.Println("Current Time : ", time.Now())
}
