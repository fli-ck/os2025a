package main

import (
	"fmt"
	"log"

	//"week10/pkg/keyboard"
	"github.com/headfirstgo/keyboard" // go get github.com/headfirstgo/keyboard
)

func main() {
	fmt.Print("실수 입력 : ")
	n, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.1f\n", n)
}
