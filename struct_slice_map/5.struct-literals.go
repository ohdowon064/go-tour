package main

import "fmt"

type Vertex5 struct {
	X, Y int
}

var (
	v1 = Vertex5{1, 2}
	v2 = Vertex5{X: 1}
	v3 = Vertex5{}
	p  = &Vertex5{1, 2}
)

func main() {
	fmt.Println(v1, v2, v3, p)
}
