package main

import "fmt"

func main() {
	// [n]T는 n개의 T 타입을 갖는 배열을 의미한다.
	var a [2]string
	a[0] = "hello"
	a[1] = "world"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
