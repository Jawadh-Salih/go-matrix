package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println("I am the legend Matrix calculator!")

	// create a matrix of size n and m
	matA := [][]int{
		{1, 2},
		{3, 4},
		{3, 4},
		{3, 4},
		{3, 4},
		{3, 4},
		{3, 4},
		{3, 4},
	}

	matB := [][]int{
		{5, 6, 1, 2, 3}, // This is a row.
		{7, 8, 3, 2, 2},
	}

	matC, err := Mul(matA, matB)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	Print(matC)
}

func Print(mat [][]int) {

	for i := range mat {
		for j := range mat[i] {
			fmt.Printf("%3d", mat[i][j])
		}
		fmt.Println()
	}
}

func Add(matA, matB [][]int) ([][]int, error) {
	if len(matA) != len(matB) {
		return nil, errors.New("Can't add arrar of length invalid")
	}

	if len(matA[0]) != len(matB[0]) {
		return nil, errors.New("Can't add arrar of length invalid")
	}

	for i := range matA {
		for j := range matA[i] {
			matA[i][j] = matA[i][j] + matB[i][j]
		}
	}

	return matA, nil
}

func Mul(A, B [][]int) ([][]int, error) {
	rowA, colA := len(A), len(A[0])
	rowB, colB := len(B), len(B[0])

	if colA != rowB {
		panic("Mismatching dimensions")
	}

	C := [][]int{}

	// loop through the rows of A (But we actuallt loop throught the rows of C)
	for r := 0; r < rowA; r++ {
		// loop through the cols of B (But we actuallr loop through the cols of C)
		arr := make([]int, 0)
		for c := 0; c < colB; c++ {
			val := 0
			// loop the column of A and multiplr row to col combination of A and B
			for j := 0; j < colA; j++ {
				// colA should be equal to rowB
				// Thus
				val += A[r][j] * B[j][c] // this gives the C[0][0] -> SUM(A[0][j]*B[j][0])
				// how to get C[0][1] -> SUM(A[0][j] * B[j][1])
				// therefore C[r][c] -> SUM of A[r][j] * B[j][c]
			}
			arr = append(arr, val)
		}
		C = append(C, arr)
	}

	return C, nil

}
