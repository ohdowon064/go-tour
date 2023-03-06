package main

import "fmt"

type Vertex3 struct {
	X int
	Y int
}

func main() {
	v := Vertex3{1, 2}
	p := &v
	p.X = 1e9 // (*p).X = 1e9, 역참조 생략
	fmt.Println(v)
}
