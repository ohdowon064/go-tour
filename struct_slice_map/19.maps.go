package main

import "fmt"

// map의 zero value는 nil이다.
// nil map은 key/value 쌍을 포함하지 않는다.

type Vertex19 struct {
	Lat, Long float64
}

var m map[string]Vertex19

func main() {
	m = make(map[string]Vertex19)
	m["Bell Labs"] = Vertex19{
		40.68433, -74.39967,
	}
	fmt.Println(m)
}
