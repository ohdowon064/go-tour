package main

import (
	"fmt"
	"math"
)

type Vertex8 struct {
	X, Y float64
}

func (v *Vertex8) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex8) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	/*
		포인트 리시버 사용 이유
		1. 메서드가 리시버의 값을 수정할 수 있다.
		2. 각 메서드 콜에서 value 복사 문제를 피할 수 있다. -> 큰 구조체라면 더욱 효율적이다.
	*/
	v := &Vertex8{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())

	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
