// Every package must have main
// Go orgonizes code based on packages not files
// so you can split code in mutiple files with
// same package name and go will recogonize
// evrything as a single package
package main

// Import is used to include packages from other libs
// Notice packages are seperated without comma
import (
	"fmt"
	"os"
	"strings"
)

// every programe execution starts with main
// init() : executes before main so you can use it
// to configure env variables or some other task
func main() {
	// := is "short variable declaration" syntax
	// go automatically recoganize datatype based
	// on value asigned to it, we can not assign
	// int value in variable bcoz its type will
	// become string once we initilize it
	who := "World"

	// Will print slice of all command line args
	// first elem of os.Args is a filename
	fmt.Println(os.Args)

	// len() is in-built method to calculate slice length
	// if user pass some value from command line
	// this condition will execute
	if len(os.Args) > 1 {

		// Join combines slice of strings
		// os.Args[1:] will skip 0th elem and start
		// with 1st elem till the end of the slice
		who = strings.Join(os.Args[1:], " ")
	}

	// Use to print value
	fmt.Println("Hello", who)
}
