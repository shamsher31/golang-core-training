package main

import (
	//"errors"
	"fmt"
)

// We can define our custome data type using
// "type" keyword, here we define empty slice of interface.
// interface{} is called empty inteface which stores
// any type of value including string, int, float, slice etc
type Stack []interface{}

// we have declare our Stack type as a slice
// so we should provide len and cap methods
// Len() : Function name
// int : return type
// (stack Stack) : called reciver with stack type
func (stack Stack) Len() int {
	return len(stack)
}

// It returns cap(capacity) of stack
// notifice here we dont pass *Stack
// because we dont want to manipulat original
// stack we just need Cap
func (stack Stack) Cap() int {
	return cap(stack)
}

// IsEmpty check elements len in stack
func (stack Stack) IsEmpty() bool {
	if len(stack) == 0 {
		return true
	}
	return false
}

// It takes pointer to stack because we want to push
// elements in stack and we are directly dealing
// with stack values. It takes x as a argument which
// is of type empty interface (inteface{}) which means
// it can accept any type of values, finally we
// append value that user wants to push in Stack interface
func (stack *Stack) Push(x interface{}) {

	// append takes slice and one or more value and returns
	// new slice with original slice content's plus given
	// value you want to add in slice
	*stack = append(*stack, x)
}

func main() {

	var myStack Stack

	fmt.Println("Stack is empty : ", myStack.IsEmpty())
	fmt.Println("Push some elements of different types in Stack")

	myStack.Push(10)
	myStack.Push([]string{"Shamsher", "Ansari"})
	myStack.Push(13.5)

	fmt.Println("Elements of stack : ", myStack)

	fmt.Println("Stack is empty : ", myStack.IsEmpty())

	if myStack.IsEmpty() == false {
		fmt.Println("Total elements in stack : ", myStack.Len())
	}

}
