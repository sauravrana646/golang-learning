package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string = "https://lco.dev"

func main() {
	fmt.Println("Welcome to Golang Web Requests")
	fmt.Println("")

	response, err := http.Get(url)
	errNotNil(err)

	defer response.Body.Close() // caller responsibility to close the connection

	byteres, err := ioutil.ReadAll(response.Body)
	errNotNil(err)
	fmt.Println("Byte Response is ", byteres)
	fmt.Println("")
	fmt.Println("String Response after conversion is ", string(byteres))
}

func errNotNil(err error) {
	if err != nil {
		panic(err)
	}
}
