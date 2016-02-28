package main

import (
	"fmt"
	"github.com/shamsher31/gostack"
)

func main() {
	// Declares new stack
	var myStack stack.Stack

	fmt.Println("Stack is empty : ", myStack.IsEmpty())

	//Push some elements of different types in Stack
	myStack.Push(10)
	myStack.Push([]string{"Shamsher", "Ansari"})
	myStack.Push(13.5)
	myStack.Push("My Awesome stack")

	fmt.Println("Elements of stack : ", myStack)

	fmt.Println("Stack is empty : ", myStack.IsEmpty())

	if myStack.IsEmpty() == false {
		fmt.Println("Total elements in stack : ", myStack.Len())
	}

	top, err := myStack.Top()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Top of the stack : ", top)

	elem, err := myStack.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Poped element : ", elem)
	fmt.Println("Elements of stack : ", myStack)
}
