package main

import "fmt"

func main() {
	fmt.Println("Welcome to Group C's Week 4 Project!")
	var s = "Golang"
	if newWord, length, err := RepeatWord(s, 5); err == nil { // Note that 'RepeatWord' is now in a different package
		fmt.Printf("Repeated word [%s], length is [%d]\n", newWord, length)
	}
	//created by Vedant Acharya(500229860)
	Pyramid(7)

	Manav()
	//created by Manav Gorasiya(500228955)

	Keyur(21, 13)
	//created by Keyur Ambekar(500229862)

	//created by Arpan Vaghani(500221873)
	getSquare(9)

	//created by Aman Sinha(500219257)

}
