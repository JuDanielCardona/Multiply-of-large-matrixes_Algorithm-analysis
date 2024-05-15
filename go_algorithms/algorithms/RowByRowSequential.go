package algorithms

import (
	"fmt"
	"time"
)

func RowByRowSequential(B, C [][]int) ([][]int, int64) {
	fmt.Println("Método: RowByRowSequential")

	startTime := time.Now().UnixNano()
	size := len(C)
	A := make([][]int, size)
	for i := range A {
		A[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				A[i][k] += B[i][j] * C[j][k]
			}
		}
	}

	endTime := time.Now().UnixNano()
	executionTime := endTime - startTime

	return A, executionTime
}
