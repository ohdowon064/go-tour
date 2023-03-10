package main

import "fmt"

func main() {
	var i interface{} = "hello"
	describe15(i)

	s := i.(string) // type assertion
	describe15(s)

	s, ok := i.(string) // type assertion
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)

}

func describe15(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
