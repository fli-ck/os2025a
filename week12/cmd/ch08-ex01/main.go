package main

import "fmt"

type subscriber struct {
	name  string
	price int
}

func applyPrice(s *subscriber) {
	s.price = 10000
	s.name = "Park Inha"
}

func main() {
	var s1 subscriber
	var p *subscriber = &s1
	//s1.name = "Kim Inha"
	applyPrice(&s1)                // 포인터를 넘겨줌
	fmt.Println(s1.name, s1.price) // Park Inha 10000
	// fmt.Println(*p.price)
	fmt.Println((*p).price) // 10000, 결국 *p.price는 s1을 가리키는 것과 같기 때문임.
	fmt.Println(p.price)    // 10000
}
