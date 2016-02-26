package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	file, err := os.Open("../.gitignore")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	stat, err := file.Stat()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Stat of file: ", stat)
	fmt.Println("Size of file: ", stat.Size(), "bytes")

	buff := make([]byte, stat.Size())

	_, err = file.Read(buff)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buff))

	// Easy way of reading file
	bufstr, err := ioutil.ReadFile("../.gitignore")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println(string(bufstr))

}
