package algorithms

import (
	"fmt"
	"time"
)

func StrassenNaiv(matrixA, matrixB [][]int) ([][]int, int64) {

	fmt.Println("Método: StrassenNaiv")

	startTime := time.Now().UnixNano()
	result := multiplication(matrixA, matrixB)
	endTime := time.Now().UnixNano()
	executionTime := endTime - startTime
	return result, executionTime
}

func addMatrix(matrixA, matrixB, matrixC [][]int, splitIndex int) {
	for i := 0; i < splitIndex; i++ {
		for j := 0; j < splitIndex; j++ {
			matrixC[i][j] = matrixA[i][j] + matrixB[i][j]
		}
	}
}

func multiplication(matrixA, matrixB [][]int) [][]int {

	col1 := len(matrixA[0])
	row1 := len(matrixA)
	col2 := len(matrixB[0])
	row2 := len(matrixB)

	if col1 != row2 {
		return [][]int{{0}}
	}

	resultMatrix := make([][]int, row1)
	for i := range resultMatrix {
		resultMatrix[i] = make([]int, col2)
	}

	if col1 == 1 {
		resultMatrix[0][0] = matrixA[0][0] * matrixB[0][0]
	} else {
		splitIndex := col1 / 2

		resultMatrix00 := make([][]int, splitIndex)
		resultMatrix01 := make([][]int, splitIndex)
		resultMatrix10 := make([][]int, splitIndex)
		resultMatrix11 := make([][]int, splitIndex)
		for i := range resultMatrix00 {
			resultMatrix00[i] = make([]int, splitIndex)
			resultMatrix01[i] = make([]int, splitIndex)
			resultMatrix10[i] = make([]int, splitIndex)
			resultMatrix11[i] = make([]int, splitIndex)
		}

		a00 := make([][]int, splitIndex)
		a01 := make([][]int, splitIndex)
		a10 := make([][]int, splitIndex)
		a11 := make([][]int, splitIndex)
		b00 := make([][]int, splitIndex)
		b01 := make([][]int, splitIndex)
		b10 := make([][]int, splitIndex)
		b11 := make([][]int, splitIndex)
		for i := range a00 {
			a00[i] = make([]int, splitIndex)
			a01[i] = make([]int, splitIndex)
			a10[i] = make([]int, splitIndex)
			a11[i] = make([]int, splitIndex)
			b00[i] = make([]int, splitIndex)
			b01[i] = make([]int, splitIndex)
			b10[i] = make([]int, splitIndex)
			b11[i] = make([]int, splitIndex)
		}

		for i := 0; i < splitIndex; i++ {
			for j := 0; j < splitIndex; j++ {
				a00[i][j] = matrixA[i][j]
				a01[i][j] = matrixA[i][j+splitIndex]
				a10[i][j] = matrixA[i+splitIndex][j]
				a11[i][j] = matrixA[i+splitIndex][j+splitIndex]
				b00[i][j] = matrixB[i][j]
				b01[i][j] = matrixB[i][j+splitIndex]
				b10[i][j] = matrixB[i+splitIndex][j]
				b11[i][j] = matrixB[i+splitIndex][j+splitIndex]
			}
		}

		addMatrix(multiplication(a00, b00), multiplication(a01, b10), resultMatrix00, splitIndex)
		addMatrix(multiplication(a00, b01), multiplication(a01, b11), resultMatrix01, splitIndex)
		addMatrix(multiplication(a10, b00), multiplication(a11, b10), resultMatrix10, splitIndex)
		addMatrix(multiplication(a10, b01), multiplication(a11, b11), resultMatrix11, splitIndex)

		for i := 0; i < splitIndex; i++ {
			for j := 0; j < splitIndex; j++ {
				resultMatrix[i][j] = resultMatrix00[i][j]
				resultMatrix[i][j+splitIndex] = resultMatrix01[i][j]
				resultMatrix[i+splitIndex][j] = resultMatrix10[i][j]
				resultMatrix[i+splitIndex][j+splitIndex] = resultMatrix11[i][j]
			}
		}
	}
	return resultMatrix
}
