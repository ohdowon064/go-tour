package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		// 클로저는 함수의 외부로부터 오는 변수를 참조하는 함수
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
