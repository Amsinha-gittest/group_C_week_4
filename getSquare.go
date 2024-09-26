package main

import "fmt"

func getSquare(a int) int {
	square := a * a
	fmt.Printf("Square of %d is %d\n", a, square)
	return square
}
