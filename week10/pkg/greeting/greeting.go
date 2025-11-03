package greeting

import "fmt"

func Hello() {
	fmt.Println("Hello!")
}

//func hi() {  // 함수 이름이 소문자로 시작할 시 외부에 노출이 안됨
func Hi() {
	fmt.Println("Hi~")
}
