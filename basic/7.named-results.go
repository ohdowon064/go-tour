package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return: named 결과를 반환한다. 가독성을 위해 짧은 함수에서만 권장
}

func main() {
	fmt.Println(split(17))
}
