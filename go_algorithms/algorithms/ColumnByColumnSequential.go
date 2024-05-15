package algorithms

import (
	"fmt"
	"time"
)

func ColumnByColumnSequential(A, B [][]int) ([][]int, int64) {

	fmt.Println("Método: ColumnByColumnSequential")

	startTime := time.Now().UnixNano()

	rowsA, colsA, colsB := len(A), len(A[0]), len(B[0])

	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
	}

	for j := 0; j < colsB; j++ {
		for i := 0; i < rowsA; i++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	endTime := time.Now().UnixNano()
	executionTime := endTime - startTime
	return result, executionTime
}
