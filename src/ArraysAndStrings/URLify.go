package main 

import (
	"fmt"
	"strings"
)

/*
URLify: Write a method to replace all spaces in a string with '%2e: You may assume that the string
has sufficient space at the end to hold the additional characters, and that you are given the "true"
length of the string. (Note: if implementing in Java, please use a character array so that you can
perform this operation in place.)
EXAMPLE
Input: "Mr John Smith      , 13
Output: "Mr%20John%20Smith"
*/

func URLify(str string, truLen int) string {
	// Count spaces
	in := strings.Split(str, "")
	c := 0
	for i:=0; i<truLen; i++ {
		if in[i] == " " {
			c++
		}
	}
	// Here c*2 because original space adds 1 
	index := truLen + c*2
	if truLen < len(in) {
		//in[truLen] = "/0"
	}
	for i:=truLen-1; i>=0; i-- {
		if in[i] == " " {
			in[index - 1] = "0"
			in[index - 2] = "2"
			in[index - 3] = "%"
			index = index - 3
		} else {
			in[index-1] = in[i]
			index--
		}
	}
	
	return strings.Join(in, "")
}

func main() {
	fmt.Println(URLify("Mr John Smith      ", 13))
}

