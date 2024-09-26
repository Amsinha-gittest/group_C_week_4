package main

import (
	"fmt"
	"strings"
)

func RepeatWord(s string, times int) (string, int, error) {
	if len(s) <= 0 {
		return "", len(s), fmt.Errorf("length of string can't be less than or equal to zero")
	}
	repeatedWord := strings.Repeat(s, times)
	fmt.Println(repeatedWord)
	return repeatedWord, len(s), nil
}
