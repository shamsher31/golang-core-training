package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var bigDigits = [][]string{
	{"   000 ",
		"0     0",
		"0     0",
		"0     0",
		"0     0",
		"0     0",
		"  000 "},
	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
	{" 222 ", "22", "2 ", " 2 ", " 2", "2", "22222"},
}

func main() {

	// if user forget to pass number show them message how to run this code
	if len(os.Args) == 1 {

		// get the last part of the filename from  0th slice
		// filepath.Base will return the filename
		fmt.Printf("Usage: %s.go <enter-number-here> \n", filepath.Base(os.Args[0]))

		// This will terminate a programe
		// 0 = Success
		// Non Zero means user error or failure
		os.Exit(1)
	}

	// Go uses single qoute for character literals and character literals
	// is an integer that's compatible with any of Go's integer type
	// Single code will return the byte(uint8) slice or code point from
	// character literal. You can get the integer value of
	// character literal by subtratcing it
	// eg: '3' - '0' = 51 - 48 = 3
	// fmt.Println('0') // 48
	// fmt.Println('A') // 65
	// fmt.Println('a') // 97

	stringOfDigits := os.Args[1]

	for row := range bigDigits {
		line := " "
		for column := range stringOfDigits {

			digit := stringOfDigits[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[column][row] + " "
			} else {

				// log.Fatal will print date, time and message on os.Stderr if
				// no other log destination is provided and call os.Exit(1)
				log.Fatal("Please enter valid number")
			}
		}
		fmt.Println(line)
	}
}
