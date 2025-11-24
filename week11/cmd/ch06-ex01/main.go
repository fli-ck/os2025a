package main

import "fmt"

func main() {
	subjects := [4]string{"Go", "Javascript", "Python", "Linux"}
	subjectsSlice := subjects[:3]
	// subjects[0] = "Java"
	subjectsSlice[0] = "Java"
	subjectsSlice = append(subjectsSlice, "Go", "Kotlin", "Database")
	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("==========")
	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}
}
