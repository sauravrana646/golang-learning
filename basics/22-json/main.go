package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Golang Json")
	fmt.Println("")

	// EncodeJson()
	DecodeJson()

}

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`                       // "-" tells to dont include this field in json
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"Mern Bootcamp", 199, "LearnCodeOnline.in", "efg123", []string{"full-stack", "reactjs"}},
		{"Golang Bootcamp", 399, "LearnCodeOnline.in", "hij123", nil},
	}

	// finalJson, _ := json.Marshal(lcoCourses)
	finalJson, _ := json.MarshalIndent(lcoCourses, "", "\t")

	fmt.Printf("Final Json is : %s\n", finalJson)

}

func DecodeJson()  {

	// var lcoCourse course

	jsonDataFromWeb := []byte(`
		{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
		}
	`)

	// if json.Valid(jsonDataFromWeb){
	// 	fmt.Println("Json was valid")

	// 	json.Unmarshal(jsonDataFromWeb, &lcoCourse)
	// 	fmt.Printf("The Data from web is : %#v\n",lcoCourse)
	// 	} else {
	// 		fmt.Println("Json was not VALID !!")
	// 	}
		
		
		// some cases where you want to add data in key value (ex : unstructured data and not known keys beforehand)
		// https://www.sohamkamani.com/golang/json/#validating-json-data
		
		var myOnlineData map[string]interface{}
		
		json.Unmarshal(jsonDataFromWeb, &myOnlineData)
		// fmt.Printf("The Online Data from web is : %#v\n",myOnlineData)

		for key,value := range myOnlineData {
			fmt.Printf("Key is : %v , Value is : %v , Type is : %T\n",key,value,value)
		}
}