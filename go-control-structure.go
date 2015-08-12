package main

import "fmt"

func main() {

	/**
	 * Go only has a one loop ie. for loop
	 * which can be used in different of ways
	 */

	for i := 1; i <= 10; i++ {

		if i%2 == 0 {
			fmt.Println(i, " even")
		} else {
			fmt.Println(i, " odd")
		}

	}

	for j := 1; j <= 5; j++ {

		switch j {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		}

	}

}
