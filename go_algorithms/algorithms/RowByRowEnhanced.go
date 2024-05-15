package algorithms

import (
	"fmt"
	"sync"
	"time"
)

func RowByRowEnhanced(B, C [][]int, bsize int) ([][]int, int64) {
	fmt.Println("Método: RowByRowEnhanced")

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
		for i1 := 0; i1 < size/2; i1 += bsize {
			for j1 := 0; j1 < size; j1 += bsize {
				for k1 := 0; k1 < size; k1 += bsize {
					for i := i1; i < i1+bsize && i < size; i++ {
						for j := j1; j < j1+bsize && j < size; j++ {
							for k := k1; k < k1+bsize && k < size; k++ {
								A[i][k] += B[i][j] * C[j][k]
							}
						}
					}
				}
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i1 := size / 2; i1 < size; i1 += bsize {
			for j1 := 0; j1 < size; j1 += bsize {
				for k1 := 0; k1 < size; k1 += bsize {
					for i := i1; i < i1+bsize && i < size; i++ {
						for j := j1; j < j1+bsize && j < size; j++ {
							for k := k1; k < k1+bsize && k < size; k++ {
								A[i][k] += B[i][j] * C[j][k]
							}
						}
					}
				}
			}
		}
	}()

	wg.Wait()

	endTime := time.Now().UnixNano()
	executionTime := endTime - startTime
	return A, executionTime
}
