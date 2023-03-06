package main

import "time"

func main() {
	t := time.Now()
	switch { // switch true와 동일, if-else 체인을 깔끔하게 가능
	case t.Hour() < 12:
		println("Good morning!")
	case t.Hour() < 17:
		println("Good afternoon.")
	default:
		println("Good evening.")
	}
}
