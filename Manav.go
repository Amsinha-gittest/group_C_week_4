package main

import "fmt"

func Manav() {
	defer fmt.Println("This will be printed last.")

	fmt.Println("First statement.")
	fmt.Println("Second statement.")
}
