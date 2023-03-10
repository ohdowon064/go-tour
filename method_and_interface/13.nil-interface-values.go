package main

import "fmt"

type I13 interface {
	M()
}

func main() {
	var i I13
	describe13(i)
	i.M() // interface tuple 내부에 타입이 없다. 따라서 i.M()은 panic을 발생시킨다.
}

func describe13(i I13) {
	fmt.Printf("(%v, %T)\n", i, i)
}
