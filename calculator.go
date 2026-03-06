package main

import "fmt"

func Divide(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func main() {
	fmt.Println(Divide(10, 2))
	fmt.Println(Divide(10, 0)) // BUG: division by zero panic
}
