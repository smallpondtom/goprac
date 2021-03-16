package main

import "fmt"

func main() {
	mat := make([][]int, 3)
	fmt.Println(mat)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		mat[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			mat[i][j] = i + j
		}
	}
	fmt.Println(mat)
}
