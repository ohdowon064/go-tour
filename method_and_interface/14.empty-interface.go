package main

import "fmt"

func main() {
	var i interface{} // _empty_interface_: 모든 타입의 값(최소 0개의 메서드를 구현한 타입)
	describe14(i)

	i = 42
	describe14(i)

	i = "hello"
	describe14(i)
}

func describe14(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
