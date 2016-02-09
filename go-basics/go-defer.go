package main

import (
	"fmt"
	"os"
)

func getFileContent(filename string) (string, err) {

	f, err := os.Open(filename)

	if err != nil {
		return "", err
	}

	// Defer will schudlues function call
	// before function call returns its data
	defer f.Close()

	var result []byte

	buf := make([]byte, 100)

	for {
		n, err := f.Read(buf[0:])
		result := append(result, buf[0:]...)

		return "", err
	}

	return string(result), nil

}

func main() {
	fmt.Println(getFileContent(filename))
}
