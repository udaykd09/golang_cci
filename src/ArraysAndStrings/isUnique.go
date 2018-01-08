package main 

import (
	"fmt"
)

//Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you
//cannot use additional data structures?
func isUniqueUsingSpace(str string) bool {
	var charArr [128]bool
	for i := 0; i < len(str); i++ {
		if !charArr[str[i]] {
		charArr[str[i]] = true
		} else { 
		return false
		}
	}
	return true
}

func isUniqueWithoutSpace(str string) bool {
	for i := 0; i < len(str); i++ {
		for j := i+1; j < len(str); j++ {
			if str[j] == str[i] {
				return false
			}
		}
	}
	return true
}

func main() {
		fmt.Println(isUniqueUsingSpace("Helo"))
		fmt.Println(isUniqueWithoutSpace("Helo"))
}