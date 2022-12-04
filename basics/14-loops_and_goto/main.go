package main

import "fmt"

func main() {
	fmt.Println("Welcome to Golang Loops")
	fmt.Println("")

	// days := []string{"Sunday","Tuesday","Wednesday","Friday","Saturday"}

	// fmt.Println("Slice is ",days)

	// for day:=0; day < len(days);day++ {
	// 	fmt.Println("Day is ",days[day])
	// }

	// for day := range days {
	// 	fmt.Println("Day is ",days[day])
	// }
	
	// for index ,day := range days {
	// 	fmt.Printf("Day on index %v is : %v\n",index,day)
	// }
	
	// for _ ,day := range days {
	// 	fmt.Println("Day is ",day)
	// }

	rogueDay := 1

	for rogueDay < 10 {

		// if rogueDay == 5 {
		// 	goto loc
		// }
		if rogueDay == 7 {
			break
		}
		// if rogueDay == 6 {
		// 	rogueDay++
		// 	continue
		// }

		fmt.Println("rogueDay value is ", rogueDay)
		rogueDay++
	}

	// loc:                        // label for go declared here and then used in loop by goto to jump to this label
	// 	fmt.Println("This is example of goto")
	
}

