package algorithms

import (
	"fmt"
	"sync"
	"time"
)

func ColumnByColumnParallel(A, B [][]int) ([][]int, int64) {
	fmt.Println("Método: ColumnByColumnParallel")

	startTime := time.Now().UnixNano()

	rowsA, colsA, colsB := len(A), len(A[0]), len(B[0])

	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
	}

	var wg sync.WaitGroup
	wg.Add(colsA * colsB)

	for j := 0; j < colsB; j++ {
		for i := 0; i < rowsA; i++ {
			go func(row, col int) {
				defer wg.Done()
				sum := 0
				for k := 0; k < colsA; k++ {
					sum += A[i][k] * B[k][j]
				}
				result[row][col] = sum
			}(i, j)
		}
	}

	wg.Wait()

	endTime := time.Now().UnixNano()
	executionTime := endTime - startTime
	return result, executionTime
}
