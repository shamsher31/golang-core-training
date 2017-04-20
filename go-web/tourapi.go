package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
)

type Tour struct {
	Name, Price string
}

func main() {

	url := "http://services.explorecalifornia.org/json/tours.php"
	content := contentFromServer(url)
	jsonContent := toursFromJson(content)

	for _, tour := range jsonContent {
		price, _, _ := big.ParseFloat(tour.Price, 10, 2, big.ToZero)
		fmt.Printf("%s ($%.2f)\n", tour.Name, price)
	}
}

func contentFromServer(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	return string(bytes)
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 1)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()

	if err != nil {
		fmt.Println(err.Error())
	}

	var tour Tour

	for decoder.More() {
		err := decoder.Decode(&tour)
		if err != nil {
			fmt.Println(err.Error())
		}
		tours = append(tours, tour)
	}
	return tours
}
