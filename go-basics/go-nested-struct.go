// Ref : https://eager.io/blog/go-and-json/

package main

import (
	"encoding/json"
	"fmt"
)

type App struct {
	Id string `json: "id"`
}

type Org struct {
	Name string `json: "name"`
}

type AppOrg struct {
	App
	Org
}

func main() {

	data := []byte(`
    {
      "id" : "36128312",
      "name" : "Awesome Organization"
    } 
 `)

	var appOrg AppOrg

	err := json.Unmarshal(data, &appOrg)

	if err != nil {
		panic(err)
	}

	fmt.Println(appOrg.App)
	fmt.Println(appOrg.Org)
	fmt.Println(appOrg.Id)
	fmt.Println(appOrg.Name)

}
