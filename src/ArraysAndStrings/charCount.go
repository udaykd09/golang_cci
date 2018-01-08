package main 

import (
	"fmt"
)

//Check if the two strings have identical character counts.
//We can also use the definition of a permutation-two words with the same character counts-to implement
//this algorithm.

func charCount(x, y string) bool {
	var letters [128]int
	if len(x) != len(y) {
		return false
	}
	for i := 0; i < len(x); i++ {
		letters[x[i]]++
	}
	
	// Trick : char with more count in x will result other char more in y 
	for i := 0; i < len(y); i++ {
		letters[y[i]]--
		if letters[y[i]] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(charCount("aaddd", "aaddd"))
}

