package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println(randomString(10))

}

func randomString(length int) string {

	var str string

	alphabet := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPKRSTUVWXYZ0123456789")

	for i := 0; i < length; i++ {
		str += string(alphabet[rand.Intn(len(alphabet))])
	}

	return str
}
