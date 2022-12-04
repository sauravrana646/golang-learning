package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


func main() {
	// const myurl string = "http://localhost:8000/get"
	// const myurl string = "http://localhost:8000/post"
	const myurl string = "http://localhost:8000/postform"
	
	fmt.Println("Welcome to Golang Web Request Verbs")
	fmt.Println("")

	// performGetRequest(myurl)
	// anotherGetHandler(myurl)
	// performPostJsonRequest(myurl)
	performPostFormRequest(myurl)
}

func performPostFormRequest(myurl string)  {
	// fake json payload
	data := url.Values{}

	data.Add("firstname", "saurav")
	data.Add("lastname", "rana")
	data.Add("email", "saurav@go.dev")

	res, err := http.PostForm(myurl,data)
	errNotNil(err)

	defer res.Body.Close()

	byteData, _ := ioutil.ReadAll(res.Body)

	fmt.Println("Status Code : ",res.StatusCode)
	fmt.Println("Content Length : ",res.ContentLength)
	fmt.Println("")
	fmt.Println("Data recieved is : ",string(byteData))

	
}

func performPostJsonRequest(myurl string)  {
	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"golang",
			"price":"100",
			"platform":"linux"
		}
	`)

	res, err := http.Post(myurl,"application/json",requestBody)
	errNotNil(err)

	defer res.Body.Close()

	byteData, _ := ioutil.ReadAll(res.Body)

	fmt.Println("Status Code : ",res.StatusCode)
	fmt.Println("Content Length : ",res.ContentLength)
	fmt.Println("")
	fmt.Println("Data recieved is : ",string(byteData))

	
}

func performGetRequest(myurl string) {

	res, err := http.Get(myurl)
	errNotNil(err)
	defer res.Body.Close()

	byteData , err := ioutil.ReadAll(res.Body)
	errNotNil(err)

	fmt.Println("Status Code : ",res.StatusCode)
	fmt.Println("Content Length : ",res.ContentLength)
	fmt.Println("")
	fmt.Println("Data recieved is : ",string(byteData))
}

func anotherGetHandler(myurl string)  {
	res, err := http.Get(myurl)
	errNotNil(err)
	defer res.Body.Close()

	var anotherresponse strings.Builder

	byteData , err := ioutil.ReadAll(res.Body)
	errNotNil(err)

	byteCount , _ :=  anotherresponse.Write(byteData)

	fmt.Println("Status Code : ",res.StatusCode)
	fmt.Println("Content Length : ",res.ContentLength)
	fmt.Println("Content Length (byteCount): ",byteCount)
	fmt.Println("")
	fmt.Println("Data recieved is : ",anotherresponse.String())

}

func errNotNil(err error){
	if err != nil {
		panic(err)
	}
}

