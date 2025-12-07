# week09

함수 사용법

```go
func GetFloat() (float64, error) { // float64, error 타입 반환
	reader := bufio.NewReader(os.Stdin) // Stdin 입력 수신
	input, err := reader.ReadString('\n') // 첫 Delim인 \n이 나오기 이전까지의 String을 읽음, (string, error) 반환
	if err != nil {
		return 0, err // 오류면 0과 error 반환
	}
	input = strings.TrimSpace(input) // Trim Space
	number, err := strconv.ParseFloat(input, 64) // Float 파싱함. 비트 사이즈는 64.
	if err != nil {
		return 0, err // 오류면 0과 error 반환
	}
	return number, nil // 오류가 없으면 Parse한 float값과 nil 반환
}

func main() { // 메인 함수
	fmt.Print("점수 입력 : ")
	score, err := GetFloat() // GetFloat 함수 실행
	if err != nil {
		log.Fatal(err) // 오류면 Fatal 내고 종료
	}
	status := "" // status 변수 선언
	if score >= 60 { // 60점 이상이면
		status = "합격!" // 합격
	} else { // 아니면
		status = "불합격" // 불합격
	}
	fmt.Printf("%.2f점은 %s\n", score, status) // 소숫점 둘 째자리 까지 포함해서 "n점은 합격/불합격!" 텍스트 출력
}
```
