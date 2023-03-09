package main

import (
	"fmt"
	"math"
)

// 인터페이스는 메소드의 시그니처 집합으로 정의된다.

type Abser interface {
	Abs() float64
}

type MyFloat9 float64

func (f MyFloat9) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex9 struct {
	X, Y float64
}

func (v *Vertex9) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat9(-math.Sqrt2)
	v := Vertex9{3, 4}

	a = f
	fmt.Println(a.Abs())
	a = &v // a = v는 컴파일 에러. Vertex9는 Abs() 메소드를 정의하지 않았기 때문이다. *Vertex9는 Abs() 메소드를 정의했기 때문에 가능하다.
	fmt.Println(a.Abs())
}