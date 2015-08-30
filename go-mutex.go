package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Mutex : Mutual Exclusion lock
// It ensures that no 2 concurrent processes
// access the same resource at the same time
func main() {
	// Initailze empty map which will be
	// used as a state
	state := make(map[int]int)

	// sync : provides synchronization primitives
	// for Mutual Exclusion lock
	// Mutex{} : is a empty struct
	// zero value for mutex is a unlocked mutex
	mutex := &sync.Mutex{}

	// ops is a number of operation count
	var ops int64 = 0

	// run 100 goroutine for Read operation
	// it reads the state of map
	for r := 0; r < 100; r++ {
		go func() {
			total := 0

			for {

				// Intn returns non-negative randome number
				key := rand.Intn(10)

				// locks map state
				mutex.Lock()

				// Access map key and add to total
				total += state[key]

				// Unlock map state
				mutex.Unlock()

				// increment operation counter
				atomic.AddInt64(&ops, 1)

				// Gosched yields the processor,
				// allowing other goroutines to run.
				// It does not suspend the current goroutine
				runtime.Gosched()

			}
		}()
	}

	// write random data to map
	for w := 0; w < 100; w++ {
		go func() {

			key := rand.Intn(10)
			val := rand.Intn(200)

			mutex.Lock()
			state[key] = val
			mutex.Unlock()

			atomic.AddInt64(&ops, 1)

			runtime.Gosched()

		}()
	}

	// allow 10 goroutines to work on
	// state and mutex
	time.Sleep(time.Second)

	// load ops count from memory
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops: ", opsFinal)

	// print state of map
	mutex.Lock()
	fmt.Println("state: ", state)
	mutex.Unlock()

}
