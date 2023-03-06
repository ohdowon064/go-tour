package main

import "fmt"

func main() {
	// 자신을 둘러싼 함수가 종료할 때까지 실행을 지연시킨다.
	// 연기된 호출 함수의 인자는 즉시 평가되지만, 함수 본체는 둘러싼 함수의 종료 이후 실행된다.
	defer fmt.Println("world")

	fmt.Println("hello")
}
