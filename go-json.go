package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Email struct {
	To      string
	Subject string
	Body    string
}

func main() {
	email := Email{"sam@gmail", "Test go mail", "This is how we encode josn"}

	encoded, err := json.Marshal(email)

	if err != nil {
		fmt.Println(err)
		return
	}

	os.Stdout.Write(encoded)

	var decoded{}
	err = json.Unmarshal(encoded, &decoded)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("%+v", decoded)

}
