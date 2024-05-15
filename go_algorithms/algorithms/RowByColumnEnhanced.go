package algorithms

import (
	"fmt"
	"sync"
	"time"
)

func RowByColumnEnhanced(B, C [][]int) ([][]int, int64) {
	fmt.Println("Método: RowByColumnEnhanced")

	startTime := time.Now().UnixNano()
	var wg sync.WaitGroup
	size := len(C)
	A := make([][]int, size)
	for i := range A {
		A[i] = make([]int, size)
	}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < size/2; i++ {
			for j := 0; j < size; j++ {
				for k := 0; k < size; k++ {
					A[i][j] += B[i][k] * C[k][j]
				}
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := size / 2; i < size; i++ {
			for j := 0; j < size; j++ {
				for k := 0; k < size; k++ {
					A[i][j] += B[i][k] * C[k][j]
				}
			}
		}
	}()

	wg.Wait()
	endTime := time.Now().UnixNano()
	executionTime := endTime - startTime
	return A, executionTime
}
