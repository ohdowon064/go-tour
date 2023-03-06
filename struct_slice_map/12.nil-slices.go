package main

import "fmt"

func main() {
	var s []int // nil slice: len, cap 0, 기본 배열이 없다.
	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("nil!")
	}
}
