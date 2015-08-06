package main

import (
  "fmt"
)

func main() {

  /**
   * Unsigned Integer which contains only positive numbers
   * uint8, uint16, uint32, uint64
   * Integer will contain numeric values
   * int8, int16, int32, int64
   */
  fmt.Println("Addition of integer 10 + 10 is : ", 10 + 10)  

  /**
   * Go has 2 floating point types
   * float32 (single precision), float64 (double precision)
   * To deal with complex numbers go has
   * complex64, complex128
   */
  fmt.Println("Addition of decimal 8.5 + 5.5 is  : ", 8.5 + 5.5)

  /**
   * len is used to calculate string lenght
   * [] is used to access string index which start at 0
   * + is used to concatenate strings
   * ``(back ticks) allow us to write string with new line
   */
  fmt.Println(len("Data type of golang"))
  fmt.Println("Data type of golang"[3])
  fmt.Println("Data type " + "of golang")
  fmt.Println(`This
    String
    is produced
    with back ticks
    so you dont have to type
    escape sequence while using
    double qoutes`)

  /**
   * Boolens are true and false
   */
  fmt.Println(true && true)
  fmt.Println(true && false)
  fmt.Println(true || true)
  fmt.Println(true || false)
  fmt.Println(!true)
  
}
