package main

import (
	"fmt"
	"time"
)

// 고루틴: 고 런타임에 의해 관리되는 경량 쓰레드

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
