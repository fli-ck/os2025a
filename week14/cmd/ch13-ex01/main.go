package main

import (
	"fmt"
	"time"
)

func say(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, ":", i)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	go say("고루틴") // 새 고루틴에서 실행
	say("메인")     // 메인 고루틴에서 실행
}
