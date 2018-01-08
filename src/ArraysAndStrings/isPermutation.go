package main 

import (
	"fmt"
	"sort"
	"strings"
)

//Check Permutation: Given two strings, write a method to decide if one is a permutation of the other.
// 1. Sort both and compare if same

func isPermutation(x, y string) bool {
	xArr := strings.Split(x, "")
	yArr := strings.Split(y, "")
	if len(xArr) != len(yArr) {
		return false
	}
	sort.Sort(sort.StringSlice(xArr))
	sort.Sort(sort.StringSlice(yArr))
	for i := 0; 	i < len(xArr); i++ {
		if xArr[i] != yArr[i] {
			return false 
		}
	}
	return true
}

func main() {
	fmt.Println(isPermutation("asd", "ds a"))
}

