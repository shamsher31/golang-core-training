package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("ansari", 0777)
	os.MkdirAll("ansari/shamsher/shabbir", 0777)

	err := os.Remove("ansari/shamsher/shabbir")

	if err != nil {
		fmt.Println(err)
	}

	os.RemoveAll("ansari")

}
