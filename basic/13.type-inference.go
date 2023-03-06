package main

import "fmt"

func main() {
	v := 42 // change me!
	fmt.Printf("v is of type %T and value is %v\n", v, v)

	// 숫자 상수는 정확도에 따라 int, float64, complex128 중 하나로 추론됩니다.
	// i := 42 		 // int
	// f := 3.142 	 // float64
	// g := 0.867 + 0.5i // complex128
}
