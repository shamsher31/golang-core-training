
 // Package Declaration
 // Every Go programe must start with Package declaration
 // Packages are Go's way of organizing and re-using code
 
package main

 // Import keyword is used to include other packages
 // fmt(format) package is used for formatting
 
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
