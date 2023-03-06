package main

import "fmt"

func main() {
	i, j := 42, 2701

	var p *int = &i // i 주소 저장
	fmt.Println(*p) // 역참조(간접참조)로 i 값 읽기
	*p = 21         // 역참조로 i 값 변경
	fmt.Println(i)

	p = &j
	*p = *p / 37 // j = j / 37
	fmt.Println(j)
}