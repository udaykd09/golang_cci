package main 

import (
	"fmt"
	"strings"
)

/*
String Rotation: Assume you have a method i5Substring which checks ifone word is a substring
of another. Given two strings, 51 and 52, write code to check if 52 is a rotation of 51 using only one
call to isSubstring (e.g., "waterbottle" is a rotation of "erbottlewat" ).
*/

func isRotated(s1, s2 string) bool {
	return strings.Contains(s1+s1, s2)
}

func main() {
	fmt.Println(isRotated("erbottlewat", "waterbottle"))
}

