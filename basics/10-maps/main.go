package main

import "fmt"

func main() {
	fmt.Println("Welcome to Golang Maps")
	fmt.Println("")

	// Initialization
	// var mymap = make(map[string]string)

	var languages = make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("Map of all languages is ", languages)
	fmt.Println("JS is short for", languages["JS"])
	
	
	// Deletion in maps
	
	delete(languages, "RB")
	fmt.Println("Map of all languages after deletion is ", languages)
}
