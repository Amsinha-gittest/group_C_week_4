package main

import (
	"strconv"
)

/*
*
Converting Integer To Binary
*
*/
func intToBin(number int) string {
	var bin string
	for number > 0 {
		bin = strconv.Itoa(number%2) + bin
		number = number / 2
	}
	return bin
}

