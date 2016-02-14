package main

// flag is used to parse command line arguments
// http allows us to request website using url
// html provides html tokenizer and parser
// TLS allow you to scan certificate
import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/shamsher31/goweblinks"
	"net/http"
	"os"
)

func main() {

	// Parse will allow to read command line args
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Please specify start page")
		os.Exit(1)
	}

	// Read the first url from command line
	getWebPage(args[0])

}

func getWebPage(uri string) {

	// This is used to allow https secure conection
	// by skiiping the certificate verification
	// when we send request to fecth URL from website
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	// Create new http client by passing transport config
	client := http.Client{Transport: transport}

	resp, err := client.Get(uri)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	// weblinks will scan webpage and give you all URL
	links := weblinks.Get(resp.Body)

	for _, link := range links {
		fmt.Println(link)
	}

}
