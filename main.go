package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Hello, world.")
	fmt.Printf("2 + 2 = %d", add(2, 2))
}
