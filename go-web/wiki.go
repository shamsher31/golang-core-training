// wiki page consist of title and body and inter linked page refrence
package main

import (
  "fmt"
  "io/ioutil"
)

// Page struct define title and body
// Body is a bytes slice bcoz ie expected by
// io library we are using
type Page struct {
  Title string
  Body []byte
}

// save method takes pointer to page as a reciver
// save does not accept any params
// error is a return type of WriteFile
// 0600 allows read write permission to file
func (p *Page) save() error {
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
}

// loadPage takes title as a param
// and returns  multiple values page pointer
// and error if the file does not exists
// ReadFile reads the content of the file
// &page returns page pointer
func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)
  
  if err != nil {
    return nil, err
  }

  return &Page{Title: title, Body: body}, nil
}

func main() {
  p1 := &Page{Title: "FirstPage", Body: []byte("Body content goes here")}
  p1.save()
  p2 , _ := loadPage("FirstPage")
  fmt.Println(string(p2.Body))
}
