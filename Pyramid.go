package main

import "fmt"

/**
 * Function to create a pyramid patern.
 *
 * @param data
 * 			rows -Input the number of rows to draw a pyramid
 * @return
 *           None
 * @author
 *          Vedant Acharya(500229860)
 */
func Pyramid(rows int) {

	fmt.Println("\nGolang program to create pyramid and pattern")
	for row := 1; row <= rows; row++ {
		for space := 1; space <= rows-row; space++ {
			fmt.Print(" ")
		}
		for pyrmd := 0; pyrmd != (2*row - 1); pyrmd++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
