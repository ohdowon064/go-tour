package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // defer는 스택으로 쌓인다. 후입선출로 9부터 출력된다.
	}

	fmt.Println("done")
}
