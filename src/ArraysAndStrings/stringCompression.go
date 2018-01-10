package main 

import (
	"fmt"
	"strings"
	"strconv"
)

func compress(x string) string {
	s := strings.ToLower(x)
	compressedStr := ""
	count := 0
	for i:=0; i<len(s); i++ {
		count++
		if i+1>=len(s) || s[i] != s[i+1] {
			compressedStr += string(s[i]) + strconv.Itoa(count)
			count = 0
		}
	}
	if len(compressedStr) < len(s) {
		return compressedStr
	} else {
		return x
	}
}

func main() {
	fmt.Println(compress("acd"))
}

