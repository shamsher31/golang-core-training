package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
)

// display usage message on terminal
var prompt = "Enter radius and angle (in degree), eg: 5.5 180, or %s to quite."

// init() will get call before main function
// GOOS (GO Operating System) returns platform we are working on
// Sprintf : uses specifier %s to replace with provided string
func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctr+Z, Enter")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

type polar struct {
	radius float64
	θ      float64
}

type cartesian struct {
	x float64
	y float64
}

func main() {

	// make is use to initialize map, slice and channels
	// chan is used to create channels which is of type
	// plolar which is a struct, it can be of string, float64 etc
	questions := make(chan polar)

	// defer will call built-in close function before main is returned
	defer close(questions)

	answers := createSolver(questions)
	defer close(answers)

	interact(questions, answers)

}

func createSolver(questions chan polar) chan cartesian {

	answers := make(chan cartesian)

	// keyword "go" is used to create goroutine
	// with self invoking anonymous function which will
	// be called asynchronousely
	go func() {

		// this will create infinite loop that waits until
		// it receives a question. This will block its own
		// goroutine and not the other goroutine and main goroutine
		for {

			// receive question from channel which will be a
			// struct of type polar
			polarCoord := <-questions

			θ := polarCoord.θ * math.Pi / 180.0
			x := polarCoord.radius * math.Cos(θ)
			y := polarCoord.radius * math.Sin(θ)

			// send calculated result to answers channel
			answers <- cartesian{x, y}

		}

	}()

	// As variable answers is declared in this function
	// Its perfectly fine to return local variable
	// because Go will garbage collect it once this function
	// is return and answers is no longer refered.
	return answers
}

const result = "Polar radius=%.02f θ=%.02f° → Cartesian x=%.02f y=%.02f\n"

func interact(questions chan polar, answers chan cartesian) {

	// create buffer reader as we want to interact user through terminal
	reader := bufio.NewReader(os.Stdin)

	// this tell the user what to enter and how to quite
	fmt.Println(prompt)

	for {
		fmt.Println("Radius and Angle")

		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		var radius, θ float64

		if _, err := fmt.Sscanf(line, "%f %f", &radius, &θ); err != nil {
			fmt.Fprintln(os.Stderr, "invalid input")
			continue
		}

		questions <- polar{radius, θ}
		coord := <-answers

		fmt.Printf(result, radius, θ, coord.x, coord.y)

	}

	fmt.Println()
}
