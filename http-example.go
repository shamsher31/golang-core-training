package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	url := "github.com"
	response, err := http.Get("http://isitup.org/" + url)

	if err != nil {
		fmt.Println("%s", err)
		os.Exit(1)
	}

	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("%s", err)
		os.Exit(1)
	}

	fmt.Println("%s\n", string(contents))
}
