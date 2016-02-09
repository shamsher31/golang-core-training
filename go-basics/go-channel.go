package main

import (
	"fmt"
	"time"
)

// '<-'(left arrow) operator is used
// to send and recive msg between channel
// c <- means send msg
// msg := <- c means recive msg
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "Ping"
	}
}

// we can specify channel direction
// adding channel arrow on right
// will make it sender
func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "Pong"
	}
}

// adding channel arrow on left
// will make it receiver
func printer(c <-chan string) {
	for {
		// print msg recived on channel
		fmt.Println(<-c)
		time.Sleep(time.Second * 1)
	}
}

// Channel provides communication
// between goroutines and sync
// execution between them
func main() {

	// 'chan' is used to create channels
	var c chan string = make(chan string)

	// call goroutine
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)

}
