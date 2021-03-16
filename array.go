package main

import "fmt"

func main() {
	var a [5]int

	fmt.Println(a)

	// Define a matrix
	var mat [5][5]int

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			mat[i][j] = len(mat)*i + j
		}
	}
	fmt.Println(mat)
}
