package main

import "fmt"

type I12 interface {
	M()
}

type T12 struct {
	S string
}

func (t *T12) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}

	fmt.Println(t.S)
}

func describe12(i I12) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I12

	var t *T12
	i = t
	describe12(i)
	i.M()

	i = &T12{"hello"}
	describe12(i)
	i.M()
}
