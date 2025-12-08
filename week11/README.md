# week11

## ch05-ex02

`github.com/headfirestgo/datafile` 모듈을 이용해 meatWeight.txt라는 텍스트를 읽고
주차별 고기 소비량 평균을 구하는 코드

```go
package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile" // github.com/headfirstgo/datafile 모듈 Import
)

func main() {
	weights, err := datafile.GetFloats("meatWeight.txt") // meatWeight.txt라는 파일이 있어야함. 그렇지 않을 경우 오류.
	if err != nil {
		log.Fatal(err) // 파일이 없거나, 기타 오류 등.
	}
	hap := 0.0 // 합 변수 선언
	for i := range weights { // for i := 0; i < len(weights); i++ 도 가능
		hap = hap + weights[i] // 합 = 합 + 무게[i]
	}
	weeks := float64(len(weights)) // 주차 == 무게 개수
	fmt.Println("주차별 고기 소비량 평균: ", hap/weeks) // 평균값 출력
}
```

## ch06-ex01

슬라이스

Array와 다르게 슬라이스는 Dynamically Sized된 형태로 동작한다.

`[]T`는 타입 `T`에 대한 슬라이스를 의미한다.

범위 지정은 Python과 비슷하게 `:`으로 가능하다.

예를 들어, `array := [4]int{1, 2, 3, 4}`에서 1, 2, 3까지만 잘라내고 싶은 경우, `newArray := [:3]`으로 구현 가능하다.

반대로 첫 Element만 빼고 나머지 전체를 사용하고 싶은 경우 `newArray2 := [1:]` 형태로 사용하면 가능하다.

array에 내용을 더 추가하고 싶은 경우에는 `append` 내장 함수를 사용한다.

`5, 6, 7`을 추가하고 싶은 경우 `newArray3 := append(newArray, 5, 6, 7)`과 같은 방식이다.

```go
package main

import "fmt"

func main() {
	subjects := [4]string{"Go", "Javascript", "Python", "Linux"} // 4개의 string이 있는 array를 생성
	subjectsSlice := subjects[:3] // Go, Javascript, Python에 대해서만 잘라낸 slice (index상 0, 1, 2)
	// subjects[0] = "Java"
	subjectsSlice[0] = "Java" // 0번 항목을 Java로 변경
	subjectsSlice = append(subjectsSlice, "Go", "Kotlin", "Database") // Java, Javascript, Python에 Go, Kotlin, Database을 append (append 함수는 slice만 가능하고 array는 불가!!)
	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("==========")
	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}
}
```

# ch06-ex02

ch05-ex02의 모듈 의존 없이 자체 구현한 버전

GetFloats만 설명

```go
func GetFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64 // 총 3주간의 Floats를 읽을 예정이므로 변수 선언 (초기화는 아직)
	file, err := os.Open(fileName) // fileName 파일명의 파일 열기, file은 `*File` 타입을 가진다.
	if err != nil {
		return numbers, err // 파일 열기 오류나면 float64 array(비어있을거임)와 오류 반환
	}
	i := 0 // i 변수 선언
	scanner := bufio.NewScanner(file) // 파일 스캐너 생성 (NewReader와 헷갈리지 않도록 하자)
	for scanner.Scan() { // 스캔을 하는 동안
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64) // numbers array i번째에 scanner.Text()를 입력값으로 한 ParseFloat 결과값을 지정, 오류값(or nil) 지정
		if err != nil {
			return numbers, err // 오류가 있으면 반환
		}
		i++ // index 1 추가
	}
	err = file.Close() // 파일 닫기 처리
	if err != nil {
		return numbers, err // 닫기 실패시 오류 반환
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err() // 스캐너 오류 존재할 시 스캐너 오류 반환
	}
	return numbers, nil // 이외에는 numbers array 반환
}
```

# ch06-ex03

ch06-ex02의 Array 제한이 없는 버전

