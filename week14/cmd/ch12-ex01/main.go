package main

import (
	"fmt"
)

func main() {
	// Last In First Out
	defer fmt.Println("1st defer")
	defer fmt.Println("2nd defer")
	defer fmt.Println("3rd defer")
	fmt.Println("main logic")
}
