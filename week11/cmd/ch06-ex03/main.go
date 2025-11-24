package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}
func main() {
	weights, err := GetFloats("meatWeight.txt")
	if err != nil {
		log.Fatal(err)
	}
	hap := 0.0
	for i := 0; i < len(weights); i++ {
		hap = hap + weights[i]
	}
	weeks := float64(len(weights))
	fmt.Println("주차별 고기 소비량 평균: ", hap/weeks)
}
