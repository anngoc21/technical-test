package main

import "fmt"

func modifyMatrix(mat [][]int) {
	var rowFlag = false
	var colFlag = false
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if i == 0 && mat[i][j] == 1 {
				rowFlag = true
			}
			if j == 0 && mat[i][j] == 1 {
				colFlag = true
			}
			if mat[i][j] == 1 {
				mat[0][j] = 1
				mat[i][0] = 1
			}

		}
	}
	for i := 1; i < len(mat); i++ {
		for j := 1; j < len(mat[0]); j++ {
			if mat[0][j] == 1 || mat[i][0] == 1 {
				mat[i][j] = 1
			}
		}
	}
	if rowFlag == true {
		for i := 0; i < len(mat[0]); i++ {
			mat[0][i] = 1
		}
	}

	// modify first col if there was any 1
	if colFlag == true {
		for i := 0; i < len(mat); i++ {
			mat[i][0] = 1
		}
	}
}

func printMatrix(mat [][]int) {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			fmt.Print(mat[i][j], " ")
		}
		fmt.Println("")
	}
}

func main() {
	mat := [][]int{
		{0, 1, 2, 3},
		{3, 1, 2, 4},
		{1, 0, 2, 3},
		{5, 9, 2, 5},
	}
	modifyMatrix(mat)
	printMatrix(mat)
}
