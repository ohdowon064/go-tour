package main

import (
	"fmt"
	"math"
)

// A concrete type은 인터페이스가 아닌 데이터타입을 말한다.
// 인터페이스 값은 값과 콘크리트 타입의 튜플이다. (value, type)

type I11 interface {
	M()
}

type T11 struct {
	S string
}

func (t *T11) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I11

	i = &T11{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

}

func describe(i I11) {
	fmt.Printf("(%v, %T)\n", i, i)
}
