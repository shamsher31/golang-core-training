package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Making progress on more then one thing
// simultaneously is called concurrency
// goroutines are define by writing 'go'
// keyword infront of func inovocation
func num(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(n, " : ", i)
		amt := time.Duration(rand.Intn(500))
		time.Sleep(time.Millisecond * amt)
	}
}

// if you do not write anything after
// goroutine the program will exit
func main() {
	for i := 1; i < 10; i++ {
		go num(i)
	}
	var input string
	fmt.Scanln(&input)
}
