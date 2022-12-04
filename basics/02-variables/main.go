package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the variables section")

	fmt.Println("Hello User, Enter you ratings:")

	buff := bufio.NewReader(os.Stdin)
	ratings, _ := buff.ReadString('\n')

	numRatings, err := strconv.ParseInt(strings.TrimSpace(ratings), 10, 32)
	// numRatings, err := strconv.ParseInt(ratings, 10, 32)

	if err != nil {
		fmt.Printf("We have some error : %v", err)
		// panic(err)
	} else {
		numRatings += 1
		fmt.Printf("Your have entered %v of type %T", numRatings, numRatings)
	}
}
