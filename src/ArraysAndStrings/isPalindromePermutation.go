package main 

import (
	"fmt"
	"strings"
)

/*
Palindrome Permutation: Given a string, write a function to check if it is a permutation of
a palindrome. A palindrome is a word or phrase that is the same forwards and backwards. A
permutation is a rearrangement of letters. The palindrome does not need to be limited to just
dictionary words.
EXAMPLE
Input: Tact Coa
Output: True (permutations:"taco cat'; "atco cta'; etc.)
*/

//NOTE : Space and Case insensitive 

func isPalindromePermutation(s string) bool {
	x := strings.ToLower(s)
	m := make(map[byte]int)
	for i:=0;i<len(x);i++ {
			m[x[i]]++
	}
	foundOdd := false
	for _,v := range m {
		if v % 2 == 1 {
			if foundOdd {
				return false
			}
			foundOdd = true
		}
	}
	return true
}


// Faster solution is to keep countOdd as we traverse the string and return false if found 2 odds 

func main() {
	fmt.Println(isPalindromePermutation("Tactcat"))
}

