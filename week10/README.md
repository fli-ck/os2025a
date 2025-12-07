# week10

### ch04 - ex01 ~ ex05 패키지 개념

- 외부에서 참조하기 위한 함수를 선언할 때에는 함수 이름을 **대문자로** 지정
- 외부에서 패키지를 받아 오기 위해 사용하는 명령은 `go get 패키지모듈이름전체`
- `go mod tidy`
- 로컬 패키지 참조는 `루트모듈명/패키지폴더명/참조패키지이름` (파일명과는 다름)
    ```go
    package main

    import (
        "log"

        "week10/pkg/keyboard"
    )

    func main() {
        input, err := keyboard.ParseFloat() // 참조하는 패키지 모듈 명의 이름은 반드시 한 번 적어줘야함!
        if err != nil {
            log.Panicln(err)
        }
        log.Println(input)
    }
    ```

### ch05-ex01 array

array 내부를 loop 하는 경우 `for index, element := range array` 형태를 사용.

```go
package main

import "fmt"

func main() {
    arr := [3]{-1, 3, 5}

    for i, number := range numbers {
        // i = index
        // number = actual nubmer value
        fmt.Println("%d %d", i, number)
    }
}
```