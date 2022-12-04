package main

import "fmt"

func main() {
	fmt.Println("Welcome to Golang slices")
	fmt.Println("")

	// Initialization
	var fruitList = []string{}
	var ptr *[]string  = &fruitList

	fmt.Println("Length of slice before is ", len(fruitList))

	fruitList = append(fruitList, "Apple")
	fmt.Println("Address of slice is ", ptr)
	fruitList = append(fruitList, "Tomato", "Peach")

	fmt.Println("Length of slice after is ", len(fruitList))
	fmt.Println("Address of slice is ", ptr)
	
	// Import appending operation
	
	var vegList = make([]string,3)
	
	fmt.Println("Length of vegList on initialization is ", len(vegList))
	
	vegList[0] = "Potato"
	vegList[1] = "Beans"
	vegList[2] = "Cabbage"
	// vegList[3] = "Eggplant"                     // This will throw an error coz initialized length is 3 
	
	
	fmt.Println("List of vegy is ", vegList)
	fmt.Println("Length of vegList after adding items is ", len(vegList))
	
	vegList = append(vegList, "Eggplant","LadyFinger")   // This will work coz this make it allocate new memory acoording to size
	
	fmt.Println("List of vegy is ", vegList)
	fmt.Println("Length of vegList after adding more items is  ", len(vegList))
	
	// Slicing the slices
	
	// vegList  = append(vegList[2:])
	// fmt.Println("Length of vegList after slicing is  ", vegList)
	
	// Delete element from slices with an index
	
	index := 2
	vegList  = append(vegList[:index], vegList[index+1:]...)
	fmt.Println("Length of vegList after slicing is  ", vegList)
	
}
