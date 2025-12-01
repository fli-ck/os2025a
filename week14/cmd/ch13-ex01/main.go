package main

import (
	"fmt"
	"time"
)

func say(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, ":", i)
		time.Sleep(2000 * time.Millisecond)
	}
}

func main() {
	start := time.Now()
	say("고루틴") // 새 고루틴에서 실행
	say("메인")  // 메인 고루틴에서 실행

	time.Sleep(1 * time.Second)
	fmt.Println("전체 실행 시간 : ", time.Since(start))
}
