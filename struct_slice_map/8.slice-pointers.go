package main

import "fmt"

func main() {
	// 슬라이스는 데이터를 저장할 수 없다. 기본 배열의 배열포인터, 길이, 용량을 가진 영역일 뿐이다.
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

}
