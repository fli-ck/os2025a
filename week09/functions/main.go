package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func main() {
	fmt.Print("점수 입력 : ")
	score, err := GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	status := ""
	if score >= 60 {
		status = "합격!"
	} else {
		status = "불합격"
	}
	fmt.Printf("%.2f점은 %s\n", score, status)
}
