package main

import "fmt"

// slice는 동적 배열로 배열보다 더 많이 쓴다.
// []T

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
