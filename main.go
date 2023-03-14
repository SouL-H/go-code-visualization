package main

import (
	"fmt"
	demo "go-visualization/demo"
)

func main() {
	test := demo.New()
	test.Set(100)
	fmt.Println(test.Get())
	test.Set(200)
	fmt.Println(test.Get())
}
