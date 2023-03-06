package main

import "fmt"

type Vertex20 struct {
	Lat, Long float64
}

func main() {
	m := map[string]Vertex20{
		"Bell Labs": Vertex20{
			40.68433, -74.39967,
		},
		"Google": Vertex20{
			37.42202, -122.08408,
		},
	}

	fmt.Println(m)
}
