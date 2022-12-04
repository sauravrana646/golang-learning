package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/tutscss?course=ruby&paymentid=sdafsadk3e4333s"

func main() {
	fmt.Println("Welcome to Golang URLs")
	fmt.Println("")

	res, _ := url.Parse(myurl)

	// fmt.Println("URL Scheme is ", res.Scheme)
	// fmt.Println("URL Host is ", res.Host)
	// fmt.Println("URL Port is ", res.Port())
	// fmt.Println("URL Path is ", res.Path)
	// fmt.Println("URL Query is ", res.RawQuery)
	
	qparams := res.Query()
	
	// fmt.Println("URL qparams is ", qparams)

	for key , val := range qparams{
		fmt.Printf("Key is : %v and Value is : %v\n", key ,val)
	}

	partsOfURL := &url.URL{
		Scheme: "https",
		Host: "google.com",
		Path: "/search",
		RawQuery: "query=linux",
	}

	generatedURL := partsOfURL.String()

	fmt.Println("Generated URL is ",generatedURL)
}
