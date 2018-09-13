package main

import "fmt"

func main() {
	num1, num2 := 1, 2
	result := num1 + num2
	if num1 < result && num2 < result {
		fmt.Println("Property 1 holds")
	} else {
		panic("Adding broke the universe")
	}
}
