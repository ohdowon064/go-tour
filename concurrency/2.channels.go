package main

import "fmt"

// channel은 사용되기 전에 make로 생성되어야한다.

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // channel c에 sum을 보낸다.
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // channel c로부터 값을 받는다.

	go func() { c <- 5 }()
	fmt.Println(<-c)

	fmt.Println(x, y, x+y)
}
