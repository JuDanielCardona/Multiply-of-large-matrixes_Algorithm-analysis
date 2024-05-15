package algorithms

import (
	"fmt"
	"sync"
	"time"
)

func RowByColumnParallel(A, B [][]int) ([][]int, int64) {
	fmt.Println("Método: RowByColumnParallel")

	startTime := time.Now().UnixNano()

	size := len(A)
	C := make([][]int, size)
	for i := range C {
		C[i] = make([]int, size)
	}

	var wg sync.WaitGroup
	wg.Add(size * size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			go func(row, col int) {
				defer wg.Done()
				sum := 0
				for k := 0; k < size; k++ {
					sum += A[row][k] * B[k][col]
				}
				C[row][col] = sum
			}(i, j)
		}
	}

	wg.Wait()

	endTime := time.Now().UnixNano()
	executionTime := endTime - startTime
	return C, executionTime
}
