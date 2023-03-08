package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // close는 반드시 sender에서만 호출되어야 한다.
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i) // close를 통해 channel이 닫히면 range는 자동으로 종료된다.
	}
}
