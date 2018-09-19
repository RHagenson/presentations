package main

import (
	"fmt"
)

func Add(x, y int) int {
	return x + y
}

func main() {
	num1, num2 := 1, 2
	result := Add(num1, num2) // num1 + num2
	if num1 < result && num2 < result {
		fmt.Println("Property 1 holds")
	} else {
		panic("Adding broke the universe")
	}
}
