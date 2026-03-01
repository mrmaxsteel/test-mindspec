package main

import "fmt"

func Divide(a, b int) int {
	return a / b
}

func main() {
	fmt.Println(Divide(10, 2))
	fmt.Println(Divide(10, 0)) // BUG: division by zero panic
}
