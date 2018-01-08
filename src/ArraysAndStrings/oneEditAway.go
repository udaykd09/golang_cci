package main 

import (
	"fmt"
)

func oneEditAway(x, y string) bool {
	if len(x) == len(y) {
		return replace(x, y)
	} else if len(x) == len(y) - 1 {
		return insert(x, y)
	} else if len(x) - 1 == len(y) {
		return insert(y, x)
	}
	return false
}

func replace(x, y string) bool {
	foundDiff := false
	for i:=0;i<len(x);i++{
		if x[i] != y[i] {
			if foundDiff {
			return false
			}
			foundDiff = true
		}
	}
	return true
}

func insert(s1, s2 string) bool {
	i1, i2 := 0, 0
	for i1 < len(s1) && i2 < len(s2) {
		if s1[i1] != s2[i2] {
			if i1 != i2 {
				return false
			}
			i2++
		} else {
			i1++
			i2++
		}
	}
	return true
}

func main() {
	fmt.Println(oneEditAway("bale", "balqq"))
}

