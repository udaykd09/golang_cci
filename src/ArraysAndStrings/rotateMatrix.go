package main 

import (
	"fmt"
)

/*
Rotate Matrix: Given an image represented by an NxN matrix, where each pixel in the image is 4
bytes, write a method to rotate the image by 90 degrees. Can you do this in place?
*/

func rotateMatrix(x[][] int) [][]int {
	n = len(x)
	for l:=0;l<n/2;l++ {
		first := l
		last := n - 1 - l
		for i := first; i < last; i ++ {
			offset := i - first
			top := x[first][i]
			// left to top
			x[first][i] = x[last-offset][first]
			// bottom to left
			
			// right to bottom
			
			// top to right
		}
	}
}

func main() {
	fmt.Println(rotateMatrix([][]))
}

