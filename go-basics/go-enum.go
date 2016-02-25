package main

import (
	"fmt"
)

// Enum in go is created using const keyword
// the value of iota is 0(zero) for the first time
// and increment by one each time if we dont specify
// in next const value in groupe
// if you specify iota for each const you can
// calculat it as power of 2(2^0, 2^1, 2^3, 2^4)
// the const value on the same level will always be same
// You can add(+1) or subtract(-1) value from iota
// 1 << is just a representation, if you remove it
// there will not be any effect on the output
const (
	Red             = (1 << iota)
	Green           = (1 << iota)
	Blue, ColorMask = (1 << iota), (1 << iota) + 1
)

func main() {
	fmt.Println(Red, Green, Blue, ColorMask)
}
