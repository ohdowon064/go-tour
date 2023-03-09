package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// 인터페이스의 암시적 구현
// Go에는 implements 키워드없이 인터페이스의 메소드를 구현함으로써 해당 인터페이스 타입을 구현하는 것이다.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
